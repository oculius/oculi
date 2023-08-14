package authz

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/oculius/oculi/v2/common/error-extension"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type (
	rbacService struct {
		enforcer Enforcer

		resourceList []string
		resourceN    int

		actionList []string
		actionN    int
	}
)

func (r *rbacService) ListResources() []string {
	var result []string
	copy(result[:], r.resourceList)
	return result
}

func (r *rbacService) ListActions() []string {
	var result []string
	copy(result[:], r.actionList)
	return result
}

func NewRBACService(db *gorm.DB, resourceList, actionList []string) RBAC {
	model :=
		`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "root"
`
	sort.Strings(resourceList)
	sort.Strings(actionList)
	return &rbacService{
		enforcer:     newEnforcer(model, db),
		resourceList: resourceList,
		actionList:   actionList,
		resourceN:    len(resourceList),
		actionN:      len(actionList),
	}
}

func (r *rbacService) Transaction(fn TxFunction) error {
	enf := r.enforcer
	err := enf.GetAdapter().(*gormadapter.Adapter).
		Transaction(enf, func(enf casbin.IEnforcer) error {
			return fn(enf)
		})
	if err != nil {
		if casted, ok := err.(errext.Error); ok {
			return casted
		}
		return ErrAuthorizationService(err, err.Error())
	}
	return nil
}

func (r *rbacService) Enforcer() Enforcer {
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

func (r *rbacService) ValidateResource(resource string) error {
	if sort.SearchStrings(r.resourceList, resource) < r.resourceN {
		return nil
	}
	return ErrInvalidResourceName(nil, map[string]any{"name": resource})
}

func (r *rbacService) ValidateAction(action string) error {
	if sort.SearchStrings(r.actionList, action) < r.actionN {
		return nil
	}
	return ErrInvalidActionName(nil, map[string]any{"name": action})
}

func (r *rbacService) Validate(resource, action string) error {
	if err := r.ValidateAction(action); err != nil {
		return err
	}

	if err := r.ValidateResource(resource); err != nil {
		return err
	}

	return nil
}

func (r *rbacService) BulkValidate(permissions Permissions) error {
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

func (r *rbacService) ListUsersForRole(role string) (Users, error) {
	users, err := r.enforcer.GetUsersForRole(RolePrefix + role)
	if err != nil {
		return nil, ErrAuthorizationService(err, err.Error())
	}
	return clearPrefix(users, UserPrefix), nil
}

func (r *rbacService) ListRolesForUser(user string) (Roles, error) {
	roles, err := r.enforcer.GetRolesForUser(UserPrefix + user)
	if err != nil {
		return nil, ErrAuthorizationService(err, err.Error())
	}
	return clearPrefix(roles, RolePrefix), nil
}

func (r *rbacService) ListPermissionsForUser(user string) Permissions {
	perms := r.enforcer.GetPermissionsForUser(UserPrefix + user)
	return newPermissions(perms)
}

func (r *rbacService) IsUserIn(user, role string) (bool, error) {
	return wrapError(r.enforcer.HasRoleForUser(UserPrefix+user, RolePrefix+role))
}

func (r *rbacService) HasRolePermission(role, resource, action string) bool {
	return r.enforcer.HasPolicy(RolePrefix+role, ResourcePrefix+resource, ActionPrefix+action)
}

func (r *rbacService) HasUserPermission(user, resource, action string) bool {
	return r.enforcer.HasPolicy(UserPrefix+user, ResourcePrefix+resource, ActionPrefix+action)
}

func (r *rbacService) AddRoleForUser(user, role string) (bool, error) {
	return wrapError(r.enforcer.AddRoleForUser(UserPrefix+user, RolePrefix+role))
}

func (r *rbacService) DelRoleForUser(user, role string) (bool, error) {
	return wrapError(r.enforcer.DeleteRoleForUser(UserPrefix+user, RolePrefix+role))
}

func (r *rbacService) AddInheritance(parentRole, childRole string) (bool, error) {
	return wrapError(r.enforcer.AddRoleForUser(RolePrefix+childRole, RolePrefix+parentRole))
}

func (r *rbacService) DelInheritance(parentRole, childRole string) (bool, error) {
	return wrapError(r.enforcer.DeleteRoleForUser(RolePrefix+childRole, RolePrefix+parentRole))
}

func (r *rbacService) AddRolesForUser(user string, roles Roles) (bool, error) {
	result := true
	err := r.Transaction(func(enf Enforcer) error {
		for _, role := range roles {
			ok, err := r.enforcer.AddRoleForUser(UserPrefix+user, RolePrefix+role)
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

func (r *rbacService) DelRolesForUser(user string, roles Roles) (bool, error) {
	result := true
	err := r.Transaction(func(enf Enforcer) error {
		for _, role := range roles {
			ok, err := r.enforcer.DeleteRoleForUser(UserPrefix+user, RolePrefix+role)
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

func (r *rbacService) DelRole(role string) (bool, error) {
	return wrapError(r.enforcer.DeleteRole(role))
}

func (r *rbacService) DelUser(user string) (bool, error) {
	return wrapError(r.enforcer.DeleteUser(user))
}

func (r *rbacService) AddPermissionForUser(user, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicy(UserPrefix+user, ResourcePrefix+resource, ActionPrefix+action))
}

func (r *rbacService) DelPermissionForUser(user, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicy(UserPrefix+user, ResourcePrefix+resource, ActionPrefix+action))
}

func (r *rbacService) AddPermissionForRole(role, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicy(RolePrefix+role, ResourcePrefix+resource, ActionPrefix+action))
}

func (r *rbacService) DelPermissionForRole(role, resource, action string) (bool, error) {
	if err := r.Validate(resource, action); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicy(RolePrefix+role, ResourcePrefix+resource, ActionPrefix+action))
}

func (r *rbacService) AddPermissionsForUser(user string, perms Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicies(perms.translate(UserPrefix + user)))
}

func (r *rbacService) DelPermissionsForUser(user string, perms Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicies(perms.translate(UserPrefix + user)))
}

func (r *rbacService) AddPermissionsForRole(role string, perms Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.AddPolicies(perms.translate(RolePrefix + role)))
}

func (r *rbacService) DelPermissionsForRole(role string, perms Permissions) (bool, error) {
	if err := r.BulkValidate(perms); err != nil {
		return false, err
	}
	return wrapError(r.enforcer.RemovePolicies(perms.translate(RolePrefix + role)))
}
