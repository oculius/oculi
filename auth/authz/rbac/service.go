package rbac

import (
	"sort"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	authz2 "github.com/oculius/oculi/v2/auth/authz"
	"github.com/oculius/oculi/v2/common/http-error"
	"github.com/oculius/oculi/v2/utils/arrayutils"
	"gorm.io/gorm"
)

type (
	service struct {
		enforcer authz2.Enforcer

		resourceList []string
		resourceN    int

		actionList []string
		actionN    int
	}
)

func NewService(db *gorm.DB, resourceList, actionList []string) Service {
	sort.Strings(resourceList)
	sort.Strings(actionList)
	return &service{
		enforcer:     authz2.NewCasbinEnforcer(accessControlModel, db),
		resourceList: resourceList,
		actionList:   actionList,
		resourceN:    len(resourceList),
		actionN:      len(actionList),
	}
}

func (r *service) AddAction(action ...string) {
	if len(action) == 0 {
		return
	}

	r.actionList = append(r.actionList, action...)
	r.actionList, _ = arrayutils.Unique[string, string](
		r.actionList, func(s string) string {
			return strings.ToLower(s)
		})
	sort.Strings(r.actionList)
	r.actionN = len(r.actionList)
}

func (r *service) AddResource(resource ...string) {
	if len(resource) == 0 {
		return
	}

	r.resourceList = append(r.resourceList, resource...)
	r.resourceList, _ = arrayutils.Unique[string, string](
		r.resourceList, func(s string) string {
			return strings.ToLower(s)
		})
	sort.Strings(r.resourceList)
	r.resourceN = len(r.resourceList)
}

func (r *service) ListResources() []string {
	var result []string
	copy(result[:], r.resourceList)
	return result
}

func (r *service) ListActions() []string {
	var result []string
	copy(result[:], r.actionList)
	return result
}

func (r *service) Transaction(fn authz2.TxFunction) error {
	enf := r.enforcer
	err := enf.GetAdapter().(*gormadapter.Adapter).
		Transaction(enf, func(enf casbin.IEnforcer) error {
			return fn(enf)
		})
	if err != nil {
		if casted, ok := err.(httperror.HttpError); ok {
			return casted
		}
		return ErrAuthorizationService(err, err.Error())
	}
	return nil
}

func (r *service) Enforcer() authz2.Enforcer {
	return r.enforcer
}

func clearPrefix(strs []string, prefix string) []string {
	for i, each := range strs {
		strs[i] = strings.Replace(each, prefix, "", 1)
	}
	return strs
}

func wrapError(ok bool, err error) (bool, error) {
	if err != nil {
		return ok, ErrAuthorizationService(err, err.Error())
	}
	return ok, nil
}

func (r *service) ValidateResource(resource string) error {
	if idx := sort.SearchStrings(r.resourceList, resource); idx < r.resourceN && r.resourceList[idx] == resource {
		return nil
	}
	return ErrInvalidResourceName(nil, map[string]any{"name": resource})
}

func (r *service) ValidateAction(action string) error {
	if idx := sort.SearchStrings(r.actionList, action); idx < r.actionN && r.actionList[idx] == action {
		return nil
	}
	return ErrInvalidActionName(nil, map[string]any{"name": action})
}

func (r *service) Validate(resource, action string) error {
	if err := r.ValidateAction(action); err != nil {
		return err
	}

	if err := r.ValidateResource(resource); err != nil {
		return err
	}

	return nil
}

func (r *service) BulkValidate(permissions authz2.Permissions) error {
	for _, each := range permissions {
		action, resource := each.Action, each.Resource
		if err := r.ValidateAction(action); err != nil {
			return err
		}

		if err := r.ValidateResource(resource); err != nil {
			return err
		}
	}
	return nil
}

func (r *service) ListUsersForRole(role string) (authz2.Users, error) {
	users, err := r.enforcer.GetUsersForRole(RolePrefix + role)
	if err != nil {
		return nil, ErrAuthorizationService(err, err.Error())
	}
	return clearPrefix(users, authz2.UserPrefix), nil
}

func (r *service) ListRolesForUser(user string) (authz2.Roles, error) {
	roles, err := r.enforcer.GetRolesForUser(authz2.UserPrefix + user)
	if err != nil {
		return nil, ErrAuthorizationService(err, err.Error())
	}
	return clearPrefix(roles, RolePrefix), nil
}

func (r *service) ListPermissionsForUser(user string) authz2.Permissions {
	perms := r.enforcer.GetPermissionsForUser(authz2.UserPrefix + user)
	return authz2.NewPermissions(perms)
}

func (r *service) IsUserIn(user, role string) (bool, error) {
	return wrapError(r.enforcer.HasRoleForUser(authz2.UserPrefix+user, RolePrefix+role))
}

func (r *service) HasRolePermission(role, resource, action string) bool {
	return r.enforcer.HasPolicy(RolePrefix+role, authz2.ResourcePrefix+resource, authz2.ActionPrefix+action)
}

func (r *service) HasUserPermission(user, resource, action string) bool {
	return r.enforcer.HasPolicy(authz2.UserPrefix+user, authz2.ResourcePrefix+resource, authz2.ActionPrefix+action)
}

func (r *service) AddRoleForUser(user, role string) (bool, error) {
	return wrapError(r.enforcer.AddRoleForUser(authz2.UserPrefix+user, RolePrefix+role))
}

func (r *service) DelRoleForUser(user, role string) (bool, error) {
	return wrapError(r.enforcer.DeleteRoleForUser(authz2.UserPrefix+user, RolePrefix+role))
}

func (r *service) AddInheritance(parentRole, childRole string) (bool, error) {
	return wrapError(r.enforcer.AddRoleForUser(RolePrefix+childRole, RolePrefix+parentRole))
}

func (r *service) DelInheritance(parentRole, childRole string) (bool, error) {
	return wrapError(r.enforcer.DeleteRoleForUser(RolePrefix+childRole, RolePrefix+parentRole))
}

func (r *service) AddRolesForUser(user string, roles authz2.Roles) (bool, error) {
	result := true
	err := r.Transaction(func(enf authz2.Enforcer) error {
		for _, role := range roles {
			ok, err := r.enforcer.AddRoleForUser(authz2.UserPrefix+user, RolePrefix+role)
			if err != nil {
				result = false
				return err
			}
			if !ok {
				result = false
				break
			}
		}
		return nil
	})
	return result, err
}

func (r *service) DelRolesForUser(user string, roles authz2.Roles) (bool, error) {
	result := true
	err := r.Transaction(func(enf authz2.Enforcer) error {
		for _, role := range roles {
			ok, err := r.enforcer.DeleteRoleForUser(authz2.UserPrefix+user, RolePrefix+role)
			if err != nil {
				result = false
				return err
			}
			if !ok {
				result = false
				break
			}
		}
		return nil
	})
	return result, err
}

func (r *service) DelRole(role string) (bool, error) {
	return wrapError(r.enforcer.DeleteRole(role))
}

func (r *service) DelUser(user string) (bool, error) {
	return wrapError(r.enforcer.DeleteUser(user))
}

func (r *service) AddPermissionForUser(user, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicy(authz2.UserPrefix+user, authz2.ResourcePrefix+resource, authz2.ActionPrefix+action))
}

func (r *service) DelPermissionForUser(user, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicy(authz2.UserPrefix+user, authz2.ResourcePrefix+resource, authz2.ActionPrefix+action))
}

func (r *service) AddPermissionForRole(role, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicy(RolePrefix+role, authz2.ResourcePrefix+resource, authz2.ActionPrefix+action))
}

func (r *service) DelPermissionForRole(role, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicy(RolePrefix+role, authz2.ResourcePrefix+resource, authz2.ActionPrefix+action))
}

func (r *service) AddPermissionsForUser(user string, perms authz2.Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicies(flattenPerms(perms, authz2.UserPrefix+user)))
}

func (r *service) DelPermissionsForUser(user string, perms authz2.Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicies(flattenPerms(perms, authz2.UserPrefix+user)))
}

func (r *service) AddPermissionsForRole(role string, perms authz2.Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicies(flattenPerms(perms, RolePrefix+role)))
}

func (r *service) DelPermissionsForRole(role string, perms authz2.Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicies(flattenPerms(perms, RolePrefix+role)))
}

func flattenPerms(p authz2.Permissions, subject string) [][]string {
	if len(p) == 0 {
		return nil
	}
	result := make([][]string, len(p))
	for i, each := range p {
		result[i] = []string{subject, authz2.ResourcePrefix + each.Resource, authz2.ActionPrefix + each.Action}
	}
	return result
}
