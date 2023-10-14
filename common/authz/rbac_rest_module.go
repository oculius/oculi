package authz

import (
	"github.com/oculius/oculi/v2/common/authn"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/server/oculi"
	"strings"
)

type (
	rbacModule struct {
		service   RBAC
		retriever UserRetrieverREST
		authn     authn.MiddlewareFactory
		setting   *RouteSetting
	}
)

func NewRBACRestModule(
	service RBAC,
	retriever UserRetrieverREST,
	auth authn.MiddlewareFactory,
	setting *RouteSetting,
) RBACRestModule {
	if setting == nil {
		setting = empty
	} else {
		setting.validate()
	}
	return &rbacModule{service, retriever, auth, setting}
}

func (r *rbacModule) Init(route oculi.RouteGroup) error {
	authReq := []oculi.MiddlewareFunc{r.authn.AuthRequired}
	root := r.setting
	route.Bundle("/authorization", func(group oculi.RouteGroup) {
		group.GET("/list", r.ListResourcesAndActions, root.ListActionsResources.middlewares(authReq, r)...)

		role := root.Role
		group.Bundle("/role", func(groupRole oculi.RouteGroup) {
			groupRole.GET("/:user", r.ListRolesForUser, role.ListRolesForUser.middlewares(authReq, r)...)
			groupRole.POST("/has-permission", r.HasRolePermission, role.HasPermission.middlewares(authReq, r)...)
			groupRole.POST("/add/inheritance", r.AddInheritance, role.AddInheritance.middlewares(authReq, r)...)
			groupRole.POST("/del/inheritance", r.DelInheritance, role.DelInheritance.middlewares(authReq, r)...)
			groupRole.POST("/del", r.DelRole, role.DelRole.middlewares(authReq, r)...)
		}, role.Root.middlewares(authReq, r)...)

		user := root.User
		group.Bundle("/user", func(groupUser oculi.RouteGroup) {
			groupUser.GET("/:user", r.ListUsersForRole, user.ListUsersForRole.middlewares(authReq, r)...)
			groupUser.POST("/check-user", r.IsUserIn, user.IsUserIn.middlewares(authReq, r)...)
			groupUser.POST("/has-permission", r.HasUserPermission, user.HasPermission.middlewares(authReq, r)...)
			groupUser.POST("/add/role", r.AddRoleForUser, user.AddRole.middlewares(authReq, r)...)
			groupUser.POST("/del/role", r.DelRoleForUser, user.DelRole.middlewares(authReq, r)...)
			groupUser.POST("/bulk-add/roles", r.AddRolesForUser, user.BulkAddRoles.middlewares(authReq, r)...)
			groupUser.POST("/bulk-del/roles", r.DelRolesForUser, user.BulkDelRoles.middlewares(authReq, r)...)
			groupUser.POST("/del", r.DelUser, user.DelUser.middlewares(authReq, r)...)
		}, user.Root.middlewares(authReq, r)...)

		perm := root.Permission
		group.Bundle("/permission", func(groupPerm oculi.RouteGroup) {
			groupPerm.GET("/:user", r.ListPermissionsForUser, perm.ListPermissionsForUser.middlewares(authReq, r)...)
			groupPerm.POST("/add/user", r.AddPermissionForUser, perm.AddPermissionForUser.middlewares(authReq, r)...)
			groupPerm.POST("/add/role", r.AddPermissionForRole, perm.AddPermissionForRole.middlewares(authReq, r)...)
			groupPerm.POST("/del/user", r.DelPermissionForUser, perm.DelPermissionForUser.middlewares(authReq, r)...)
			groupPerm.POST("/del/role", r.DelPermissionForRole, perm.DelPermissionForRole.middlewares(authReq, r)...)
			groupPerm.POST("/bulk-add/user", r.AddPermissionsForUser, perm.BulkAddPermissionsForUser.middlewares(authReq, r)...)
			groupPerm.POST("/bulk-add/role", r.AddPermissionsForRole, perm.BulkAddPermissionsForRole.middlewares(authReq, r)...)
			groupPerm.POST("/bulk-del/user", r.DelPermissionsForUser, perm.BulkDelPermissionsForUser.middlewares(authReq, r)...)
			groupPerm.POST("/bulk-del/role", r.DelPermissionsForRole, perm.BulkDelPermissionsForRole.middlewares(authReq, r)...)
		}, perm.Root.middlewares(authReq, r)...)
	}, root.Root.middlewares(authReq, r)...)
	route.GET("/authorization", r.ListCurrentUserPermissions, r.setting.ListCurrentUserPermissions.middlewares(authReq, r)...)
	return nil
}

func (r *rbacModule) getIdentifier(ctx oculi.Context) string {
	identifier := strings.TrimSpace(r.retriever(ctx))
	if len(identifier) == 0 {
		identifier = PublicIdentifier
	}
	return identifier
}

func (r *rbacModule) ListCurrentUserPermissions(ctx oculi.Context) error {
	identifier := r.getIdentifier(ctx)
	perms := r.service.ListPermissionsForUser(identifier)
	return ctx.AutoSend(response.NewResponse("success", map[string]any{
		"user":        identifier,
		"permissions": perms,
	}, nil))
}

func (r *rbacModule) ListResourcesAndActions(ctx oculi.Context) error {
	return ctx.AutoSend(response.NewResponse("success", map[string]any{
		"resources": r.service.ListResources(),
		"actions":   r.service.ListActions(),
	}, nil))
}

func (r *rbacModule) Permission(resource, action string) oculi.MiddlewareFunc {
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
