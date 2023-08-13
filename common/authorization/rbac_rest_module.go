package authorization

import (
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"strings"
)

type (
	rbacModule struct {
		service   RBAC
		retriever UserRetrieverREST
	}
)

func NewRBACRestModule(service RBAC, retriever UserRetrieverREST) RBACRestModule {
	return &rbacModule{service, retriever}
}

func (r *rbacModule) Init(group oculi.RouteGroup) error {
	// TODO setting
	return nil
}

func (r *rbacModule) Permission(resource, action string) oculi.MiddlewareFunc {
	return func(next oculi.HandlerFunc) oculi.HandlerFunc {
		return func(ctx oculi.Context) error {
			identifier := strings.TrimSpace(r.retriever(ctx))
			if len(identifier) == 0 {
				identifier = PublicIdentifier
			}

			hasPermission := r.service.HasUserPermission(identifier, resource, action)
			if !hasPermission {
				return ErrForbidden(nil, nil)
			}
			return next(ctx)
		}
	}
}

func (r *rbacModule) ListUsersForRole(ctx oculi.Context) error {
	return getterHelper(ctx, "role", func(data string) (any, error) {
		return r.service.ListUsersForRole(data)
	})
}

func (r *rbacModule) ListRolesForUser(ctx oculi.Context) error {
	return getterHelper(ctx, "user", func(data string) (any, error) {
		return r.service.ListRolesForUser(data)
	})
}

func (r *rbacModule) ListPermissionsForUser(ctx oculi.Context) error {
	return getterHelper(ctx, "user", func(data string) (any, error) {
		return r.service.ListPermissionsForUser(data), nil
	})
}

func (r *rbacModule) IsUserIn(ctx oculi.Context) error {
	return twinStringCheckerEndpoint[userRoleReq](r.service.IsUserIn, ctx)
}

func (r *rbacModule) HasRolePermission(ctx oculi.Context) error {
	return dataCheckerEndpoint[roleSinglePermReq](r.service.HasRolePermission, ctx)
}

func (r *rbacModule) HasUserPermission(ctx oculi.Context) error {
	return dataCheckerEndpoint[userSinglePermReq](r.service.HasUserPermission, ctx)
}

func (r *rbacModule) AddRoleForUser(ctx oculi.Context) error {
	return twinStringEndpoint[userRoleReq](r.service.AddRoleForUser, ctx)
}

func (r *rbacModule) DelRoleForUser(ctx oculi.Context) error {
	return twinStringEndpoint[userRoleReq](r.service.DelRoleForUser, ctx)
}

func (r *rbacModule) AddInheritance(ctx oculi.Context) error {
	return twinStringEndpoint[twinRoleReq](r.service.AddInheritance, ctx)
}

func (r *rbacModule) DelInheritance(ctx oculi.Context) error {
	return twinStringEndpoint[twinRoleReq](r.service.DelInheritance, ctx)
}

func (r *rbacModule) AddRolesForUser(ctx oculi.Context) error {
	return stringRolesEndpoint[userRolesReq](r.service.AddRolesForUser, ctx)
}

func (r *rbacModule) DelRolesForUser(ctx oculi.Context) error {
	return stringRolesEndpoint[userRolesReq](r.service.DelRolesForUser, ctx)
}

func (r *rbacModule) DelRole(ctx oculi.Context) error {
	return singleStringEndpoint[roleReq](r.service.DelRole, ctx)
}

func (r *rbacModule) DelUser(ctx oculi.Context) error {
	return singleStringEndpoint[userReq](r.service.DelUser, ctx)
}

func (r *rbacModule) AddPermissionForUser(ctx oculi.Context) error {
	return tripleDataEndpoint[userSinglePermReq](r.service.AddPermissionForUser, ctx)
}

func (r *rbacModule) DelPermissionForUser(ctx oculi.Context) error {
	return tripleDataEndpoint[userSinglePermReq](r.service.DelPermissionForUser, ctx)
}

func (r *rbacModule) AddPermissionForRole(ctx oculi.Context) error {
	return tripleDataEndpoint[roleSinglePermReq](r.service.AddPermissionForRole, ctx)
}

func (r *rbacModule) DelPermissionForRole(ctx oculi.Context) error {
	return tripleDataEndpoint[roleSinglePermReq](r.service.DelPermissionForRole, ctx)
}

func (r *rbacModule) AddPermissionsForUser(ctx oculi.Context) error {
	return bulkPermEndpoint[userBulkPermsReq](r.service.AddPermissionsForUser, ctx)
}

func (r *rbacModule) DelPermissionsForUser(ctx oculi.Context) error {
	return bulkPermEndpoint[userBulkPermsReq](r.service.DelPermissionsForUser, ctx)
}

func (r *rbacModule) AddPermissionsForRole(ctx oculi.Context) error {
	return bulkPermEndpoint[roleBulkPermsReq](r.service.AddPermissionsForRole, ctx)
}

func (r *rbacModule) DelPermissionsForRole(ctx oculi.Context) error {
	return bulkPermEndpoint[roleBulkPermsReq](r.service.DelPermissionsForRole, ctx)
}
