package authz

import (
	errext "github.com/oculius/oculi/v2/common/error-extension"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"net/http"
)

type (
	RBAC interface {
		Service

		ListUsersForRole(role string) (Users, error)
		ListRolesForUser(user string) (Roles, error)
		ListPermissionsForUser(user string) Permissions

		ListResources() []string
		ListActions() []string

		ValidateResource(resource string) error
		ValidateAction(action string) error
		Validate(resource, action string) error
		BulkValidate(permissions Permissions) error

		IsUserIn(user, role string) (bool, error)
		HasRolePermission(role, resource, action string) bool
		HasUserPermission(user, resource, action string) bool

		AddRoleForUser(user, role string) (bool, error)
		DelRoleForUser(user, role string) (bool, error)
		AddInheritance(parentRole, childRole string) (bool, error)
		DelInheritance(parentRole, childRole string) (bool, error)

		AddRolesForUser(user string, roles Roles) (bool, error)
		DelRolesForUser(user string, roles Roles) (bool, error)

		DelRole(role string) (bool, error)
		DelUser(user string) (bool, error)

		AddPermissionForUser(user, resource, action string) (bool, error)
		DelPermissionForUser(user, resource, action string) (bool, error)
		AddPermissionForRole(role, resource, action string) (bool, error)
		DelPermissionForRole(role, resource, action string) (bool, error)

		AddPermissionsForUser(user string, perms Permissions) (bool, error)
		DelPermissionsForUser(user string, perms Permissions) (bool, error)
		AddPermissionsForRole(role string, perms Permissions) (bool, error)
		DelPermissionsForRole(role string, perms Permissions) (bool, error)
	}

	RBACRestModule interface {
		rest.Module

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

var (
	PublicIdentifier = "guest"

	ErrAuthorizationService = errext.New("authorization service error", http.StatusInternalServerError)
	ErrInvalidResourceName  = errext.New("invalid resource name", http.StatusBadRequest)
	ErrInvalidActionName    = errext.New("invalid action name", http.StatusBadRequest)
	ErrForbidden            = errext.New("no permission", http.StatusForbidden)
)

const RolePrefix = "role_"
