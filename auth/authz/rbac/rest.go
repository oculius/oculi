package rbac

import (
	"strings"

	"github.com/oculius/oculi/v2/auth/authn"
	"github.com/oculius/oculi/v2/auth/authz"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest/oculi"
)

type (
	restModule struct {
		service             Service
		retriever           authz.UserRetrieverREST
		authn               authn.MiddlewareFactory
		setting             *RouteSetting
		resourcePermissions map[string][]string
	}
)

func NewRestModule(
	service Service,
	retriever authz.UserRetrieverREST,
	auth authn.MiddlewareFactory,
	setting *RouteSetting,
) RestModule {
	if setting == nil {
		setting = emptySetting
	} else {
		setting.validate()
	}
	return &restModule{service, retriever,
		auth, setting, map[string][]string{}}
}

func (r *restModule) Init(route oculi.RouteGroup) error {
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

func (r *restModule) getIdentifier(ctx oculi.Context) string {
	identifier := strings.TrimSpace(r.retriever(ctx))
	if len(identifier) == 0 {
		identifier = authz.PublicIdentifier
	}
	return identifier
}

func (r *restModule) ListCurrentUserPermissions(ctx oculi.Context) error {
	identifier := r.getIdentifier(ctx)
	perms := r.service.ListPermissionsForUser(identifier)
	return ctx.AutoSend(response.NewResponse("success", map[string]any{
		"user":        identifier,
		"permissions": perms,
	}, nil))
}

func (r *restModule) ListResourcesAndActions(ctx oculi.Context) error {
	return ctx.AutoSend(response.NewResponse("success", map[string]any{
		"resources":                r.service.ListResources(),
		"actions":                  r.service.ListActions(),
		"resourceToActionsMapping": r.resourcePermissions,
	}, nil))
}

func (r *restModule) Permission(resource, action string) oculi.MiddlewareFunc {
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

func (r *restModule) ListUsersForRole(ctx oculi.Context) error {
	return getterHelper(ctx, "role", func(data string) (any, error) {
		return r.service.ListUsersForRole(data)
	})
}

func (r *restModule) ListRolesForUser(ctx oculi.Context) error {
	return getterHelper(ctx, "user", func(data string) (any, error) {
		return r.service.ListRolesForUser(data)
	})
}

func (r *restModule) ListPermissionsForUser(ctx oculi.Context) error {
	return getterHelper(ctx, "user", func(data string) (any, error) {
		return r.service.ListPermissionsForUser(data), nil
	})
}

func (r *restModule) IsUserIn(ctx oculi.Context) error {
	return twinStringCheckerEndpoint[userRoleReq](r.service.IsUserIn, ctx)
}

func (r *restModule) HasRolePermission(ctx oculi.Context) error {
	return dataCheckerEndpoint[roleSinglePermReq](r.service.HasRolePermission, ctx)
}

func (r *restModule) HasUserPermission(ctx oculi.Context) error {
	return dataCheckerEndpoint[userSinglePermReq](r.service.HasUserPermission, ctx)
}

func (r *restModule) AddRoleForUser(ctx oculi.Context) error {
	return twinStringEndpoint[userRoleReq](r.service.AddRoleForUser, ctx)
}

func (r *restModule) DelRoleForUser(ctx oculi.Context) error {
	return twinStringEndpoint[userRoleReq](r.service.DelRoleForUser, ctx)
}

func (r *restModule) AddInheritance(ctx oculi.Context) error {
	return twinStringEndpoint[twinRoleReq](r.service.AddInheritance, ctx)
}

func (r *restModule) DelInheritance(ctx oculi.Context) error {
	return twinStringEndpoint[twinRoleReq](r.service.DelInheritance, ctx)
}

func (r *restModule) AddRolesForUser(ctx oculi.Context) error {
	return stringRolesEndpoint[userRolesReq](r.service.AddRolesForUser, ctx)
}

func (r *restModule) DelRolesForUser(ctx oculi.Context) error {
	return stringRolesEndpoint[userRolesReq](r.service.DelRolesForUser, ctx)
}

func (r *restModule) DelRole(ctx oculi.Context) error {
	return singleStringEndpoint[roleReq](r.service.DelRole, ctx)
}

func (r *restModule) DelUser(ctx oculi.Context) error {
	return singleStringEndpoint[userReq](r.service.DelUser, ctx)
}

func (r *restModule) AddPermissionForUser(ctx oculi.Context) error {
	return tripleDataEndpoint[userSinglePermReq](r.service.AddPermissionForUser, ctx)
}

func (r *restModule) DelPermissionForUser(ctx oculi.Context) error {
	return tripleDataEndpoint[userSinglePermReq](r.service.DelPermissionForUser, ctx)
}

func (r *restModule) AddPermissionForRole(ctx oculi.Context) error {
	return tripleDataEndpoint[roleSinglePermReq](r.service.AddPermissionForRole, ctx)
}

func (r *restModule) DelPermissionForRole(ctx oculi.Context) error {
	return tripleDataEndpoint[roleSinglePermReq](r.service.DelPermissionForRole, ctx)
}

func (r *restModule) AddPermissionsForUser(ctx oculi.Context) error {
	return bulkPermEndpoint[userBulkPermsReq](r.service.AddPermissionsForUser, ctx)
}

func (r *restModule) DelPermissionsForUser(ctx oculi.Context) error {
	return bulkPermEndpoint[userBulkPermsReq](r.service.DelPermissionsForUser, ctx)
}

func (r *restModule) AddPermissionsForRole(ctx oculi.Context) error {
	return bulkPermEndpoint[roleBulkPermsReq](r.service.AddPermissionsForRole, ctx)
}

func (r *restModule) DelPermissionsForRole(ctx oculi.Context) error {
	return bulkPermEndpoint[roleBulkPermsReq](r.service.DelPermissionsForRole, ctx)
}

func (r *restModule) AddPermissionMap(resource, action string) {
	if r.resourcePermissions == nil {
		r.resourcePermissions = map[string][]string{}
	}
	perms := r.resourcePermissions[resource]
	if perms == nil {
		perms = make([]string, 10)
	}
	perms = append(perms, action)
	r.resourcePermissions[resource] = perms
}
