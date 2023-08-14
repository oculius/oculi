package authz

import (
	"github.com/oculius/oculi/v2/common/authn"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"strings"
)

type (
	rbacModule[V any] struct {
		service   RBAC
		retriever UserRetrieverREST
		authn     authn.Factory[V]
	}
)

func NewRBACRestModule[V any](
	service RBAC,
	retriever UserRetrieverREST,
	auth authn.Factory[V],
) RBACRestModule {
	return &rbacModule[V]{service, retriever, auth}
}

func (r *rbacModule[V]) Init(route oculi.RouteGroup) error {
	authReq := r.authn.AuthRequired
	route.Bundle("/authorization", func(group oculi.RouteGroup) {
		group.GET("/list", r.ListResourcesAndActions)

		group.Bundle("/role", func(groupRole oculi.RouteGroup) {
			groupRole.GET("/:user", r.ListRolesForUser)
			groupRole.POST("/has-permission", r.HasRolePermission)
			groupRole.POST("/add/inheritance", r.AddInheritance)
			groupRole.POST("/del/inheritance", r.DelInheritance)
			groupRole.POST("/del", r.DelRole)
		})
		group.Bundle("/user", func(groupUser oculi.RouteGroup) {
			groupUser.GET("/:user", r.ListUsersForRole)
			groupUser.POST("/check-user", r.IsUserIn)
			groupUser.POST("/has-permission", r.HasUserPermission)
			groupUser.POST("/add/role", r.AddRoleForUser)
			groupUser.POST("/del/role", r.DelRoleForUser)
			groupUser.POST("/bulk-add/roles", r.AddRolesForUser)
			groupUser.POST("/bulk-del/roles", r.DelRolesForUser)
			groupUser.POST("/del", r.DelUser)
		})
		group.Bundle("/permission", func(groupPerm oculi.RouteGroup) {
			groupPerm.GET("/:user", r.ListPermissionsForUser)
			groupPerm.POST("/add/user", r.AddPermissionForUser)
			groupPerm.POST("/add/role", r.AddPermissionForRole)
			groupPerm.POST("/del/user", r.DelPermissionForUser)
			groupPerm.POST("/del/role", r.DelPermissionForRole)
			groupPerm.POST("/bulk-add/user", r.AddPermissionsForUser)
			groupPerm.POST("/bulk-add/role", r.AddPermissionsForRole)
			groupPerm.POST("/bulk-del/user", r.DelPermissionsForUser)
			groupPerm.POST("/bulk-del/role", r.DelPermissionsForRole)
		})
	}, authReq)
	route.GET("/authorization", r.ListCurrentUserPermissions)
	return nil
}

func (r *rbacModule[V]) getIdentifier(ctx oculi.Context) string {
	identifier := strings.TrimSpace(r.retriever(ctx))
	if len(identifier) == 0 {
		identifier = PublicIdentifier
	}
	return identifier
}

func (r *rbacModule[V]) ListCurrentUserPermissions(ctx oculi.Context) error {
	identifier := r.getIdentifier(ctx)
	perms := r.service.ListPermissionsForUser(identifier)
	return ctx.AutoSend(response.NewResponse("success", map[string]any{
		"user":        identifier,
		"permissions": perms,
	}, nil))
}

func (r *rbacModule[V]) ListResourcesAndActions(ctx oculi.Context) error {
	return ctx.AutoSend(response.NewResponse("success", map[string]any{
		"resources": r.service.ListResources(),
		"actions":   r.service.ListActions(),
	}, nil))
}

func (r *rbacModule[V]) Permission(resource, action string) oculi.MiddlewareFunc {
	return func(next oculi.HandlerFunc) oculi.HandlerFunc {
		return func(ctx oculi.Context) error {
			identifier := r.getIdentifier(ctx)
			hasPermission := r.service.HasUserPermission(identifier, resource, action)
			if !hasPermission {
				return ErrForbidden(nil, nil)
			}
			return next(ctx)
		}
	}
}

func (r *rbacModule[V]) ListUsersForRole(ctx oculi.Context) error {
	return getterHelper(ctx, "role", func(data string) (any, error) {
		return r.service.ListUsersForRole(data)
	})
}

func (r *rbacModule[V]) ListRolesForUser(ctx oculi.Context) error {
	return getterHelper(ctx, "user", func(data string) (any, error) {
		return r.service.ListRolesForUser(data)
	})
}

func (r *rbacModule[V]) ListPermissionsForUser(ctx oculi.Context) error {
	return getterHelper(ctx, "user", func(data string) (any, error) {
		return r.service.ListPermissionsForUser(data), nil
	})
}

func (r *rbacModule[V]) IsUserIn(ctx oculi.Context) error {
	return twinStringCheckerEndpoint[userRoleReq](r.service.IsUserIn, ctx)
}

func (r *rbacModule[V]) HasRolePermission(ctx oculi.Context) error {
	return dataCheckerEndpoint[roleSinglePermReq](r.service.HasRolePermission, ctx)
}

func (r *rbacModule[V]) HasUserPermission(ctx oculi.Context) error {
	return dataCheckerEndpoint[userSinglePermReq](r.service.HasUserPermission, ctx)
}

func (r *rbacModule[V]) AddRoleForUser(ctx oculi.Context) error {
	return twinStringEndpoint[userRoleReq](r.service.AddRoleForUser, ctx)
}

func (r *rbacModule[V]) DelRoleForUser(ctx oculi.Context) error {
	return twinStringEndpoint[userRoleReq](r.service.DelRoleForUser, ctx)
}

func (r *rbacModule[V]) AddInheritance(ctx oculi.Context) error {
	return twinStringEndpoint[twinRoleReq](r.service.AddInheritance, ctx)
}

func (r *rbacModule[V]) DelInheritance(ctx oculi.Context) error {
	return twinStringEndpoint[twinRoleReq](r.service.DelInheritance, ctx)
}

func (r *rbacModule[V]) AddRolesForUser(ctx oculi.Context) error {
	return stringRolesEndpoint[userRolesReq](r.service.AddRolesForUser, ctx)
}

func (r *rbacModule[V]) DelRolesForUser(ctx oculi.Context) error {
	return stringRolesEndpoint[userRolesReq](r.service.DelRolesForUser, ctx)
}

func (r *rbacModule[V]) DelRole(ctx oculi.Context) error {
	return singleStringEndpoint[roleReq](r.service.DelRole, ctx)
}

func (r *rbacModule[V]) DelUser(ctx oculi.Context) error {
	return singleStringEndpoint[userReq](r.service.DelUser, ctx)
}

func (r *rbacModule[V]) AddPermissionForUser(ctx oculi.Context) error {
	return tripleDataEndpoint[userSinglePermReq](r.service.AddPermissionForUser, ctx)
}

func (r *rbacModule[V]) DelPermissionForUser(ctx oculi.Context) error {
	return tripleDataEndpoint[userSinglePermReq](r.service.DelPermissionForUser, ctx)
}

func (r *rbacModule[V]) AddPermissionForRole(ctx oculi.Context) error {
	return tripleDataEndpoint[roleSinglePermReq](r.service.AddPermissionForRole, ctx)
}

func (r *rbacModule[V]) DelPermissionForRole(ctx oculi.Context) error {
	return tripleDataEndpoint[roleSinglePermReq](r.service.DelPermissionForRole, ctx)
}

func (r *rbacModule[V]) AddPermissionsForUser(ctx oculi.Context) error {
	return bulkPermEndpoint[userBulkPermsReq](r.service.AddPermissionsForUser, ctx)
}

func (r *rbacModule[V]) DelPermissionsForUser(ctx oculi.Context) error {
	return bulkPermEndpoint[userBulkPermsReq](r.service.DelPermissionsForUser, ctx)
}

func (r *rbacModule[V]) AddPermissionsForRole(ctx oculi.Context) error {
	return bulkPermEndpoint[roleBulkPermsReq](r.service.AddPermissionsForRole, ctx)
}

func (r *rbacModule[V]) DelPermissionsForRole(ctx oculi.Context) error {
	return bulkPermEndpoint[roleBulkPermsReq](r.service.DelPermissionsForRole, ctx)
}
