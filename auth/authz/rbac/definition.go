package rbac

import (
	authz2 "github.com/oculius/oculi/v2/auth/authz"
	"github.com/oculius/oculi/v2/rest"
	"github.com/oculius/oculi/v2/rest/oculi"
)

const (
	accessControlModel = `
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
)

type (
	Service interface {
		authz2.CasbinService

		ListUsersForRole(role string) (authz2.Users, error)
		ListRolesForUser(user string) (authz2.Roles, error)
		ListPermissionsForUser(user string) authz2.Permissions

		ListResources() []string
		ListActions() []string

		AddAction(action ...string)
		AddResource(resource ...string)

		ValidateResource(resource string) error
		ValidateAction(action string) error
		Validate(resource, action string) error
		BulkValidate(permissions authz2.Permissions) error

		IsUserIn(user, role string) (bool, error)
		HasRolePermission(role, resource, action string) bool
		HasUserPermission(user, resource, action string) bool

		AddRoleForUser(user, role string) (bool, error)
		DelRoleForUser(user, role string) (bool, error)
		AddInheritance(parentRole, childRole string) (bool, error)
		DelInheritance(parentRole, childRole string) (bool, error)

		AddRolesForUser(user string, roles authz2.Roles) (bool, error)
		DelRolesForUser(user string, roles authz2.Roles) (bool, error)

		DelRole(role string) (bool, error)
		DelUser(user string) (bool, error)

		AddPermissionForUser(user, resource, action string) (bool, error)
		DelPermissionForUser(user, resource, action string) (bool, error)
		AddPermissionForRole(role, resource, action string) (bool, error)
		DelPermissionForRole(role, resource, action string) (bool, error)

		AddPermissionsForUser(user string, perms authz2.Permissions) (bool, error)
		DelPermissionsForUser(user string, perms authz2.Permissions) (bool, error)
		AddPermissionsForRole(role string, perms authz2.Permissions) (bool, error)
		DelPermissionsForRole(role string, perms authz2.Permissions) (bool, error)
	}

	RestModule interface {
		rest.AccessPoint

		AddPermissionMap(resource, action string)
		Permission(resource, action string) oculi.MiddlewareFunc

		ListUsersForRole(ctx oculi.Context) error
		ListRolesForUser(ctx oculi.Context) error
		ListPermissionsForUser(ctx oculi.Context) error
		ListResourcesAndActions(ctx oculi.Context) error
		ListCurrentUserPermissions(ctx oculi.Context) error

		IsUserIn(ctx oculi.Context) error
		HasRolePermission(ctx oculi.Context) error
		HasUserPermission(ctx oculi.Context) error

		AddRoleForUser(ctx oculi.Context) error
		DelRoleForUser(ctx oculi.Context) error
		AddInheritance(ctx oculi.Context) error
		DelInheritance(ctx oculi.Context) error

		AddRolesForUser(ctx oculi.Context) error
		DelRolesForUser(ctx oculi.Context) error

		DelRole(ctx oculi.Context) error
		DelUser(ctx oculi.Context) error

		AddPermissionForUser(ctx oculi.Context) error
		DelPermissionForUser(ctx oculi.Context) error
		AddPermissionForRole(ctx oculi.Context) error
		DelPermissionForRole(ctx oculi.Context) error

		AddPermissionsForUser(ctx oculi.Context) error
		DelPermissionsForUser(ctx oculi.Context) error
		AddPermissionsForRole(ctx oculi.Context) error
		DelPermissionsForRole(ctx oculi.Context) error
	}
)

const RolePrefix = "role_"
