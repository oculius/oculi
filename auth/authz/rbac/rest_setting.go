package rbac

import (
	"reflect"
	"strings"

	"github.com/oculius/oculi/v2/rest/oculi"
)

type (
	Setting struct {
		Middlewares []oculi.MiddlewareFunc
		Name        string
		Permission  string
	}

	RouteSetting struct {
		Root                       *Setting
		ListActionsResources       *Setting
		Role                       *RoleRouteSetting
		User                       *UserRouteSetting
		Permission                 *PermissionRouteSetting
		ListCurrentUserPermissions *Setting
	}

	RoleRouteSetting struct {
		Root             *Setting
		ListRolesForUser *Setting
		HasPermission    *Setting
		AddInheritance   *Setting
		DelInheritance   *Setting
		DelRole          *Setting
	}

	UserRouteSetting struct {
		Root             *Setting
		ListUsersForRole *Setting
		IsUserIn         *Setting
		HasPermission    *Setting
		AddRole          *Setting
		DelRole          *Setting
		BulkAddRoles     *Setting
		BulkDelRoles     *Setting
		DelUser          *Setting
	}

	PermissionRouteSetting struct {
		Root                      *Setting
		ListPermissionsForUser    *Setting
		AddPermissionForUser      *Setting
		AddPermissionForRole      *Setting
		DelPermissionForUser      *Setting
		DelPermissionForRole      *Setting
		BulkAddPermissionsForUser *Setting
		BulkAddPermissionsForRole *Setting
		BulkDelPermissionsForUser *Setting
		BulkDelPermissionsForRole *Setting
	}
)

func (s *Setting) ok() bool {
	if len(strings.TrimSpace(s.Name)) == 0 ||
		len(strings.TrimSpace(s.Permission)) == 0 {
		return false
	}

	return true
}

func (s *Setting) permissionChecker(r RestModule) oculi.MiddlewareFunc {
	if len(strings.TrimSpace(s.Permission)) > 0 && len(strings.TrimSpace(s.Name)) > 0 {
		return r.Permission(s.Permission, s.Name)
	}
	return nil
}

func (s *Setting) middlewares(base []oculi.MiddlewareFunc, r RestModule) []oculi.MiddlewareFunc {
	var result []oculi.MiddlewareFunc
	if len(base) > 0 {
		copy(result, base[:])
	} else {
		result = []oculi.MiddlewareFunc{}
	}
	if permChecker := s.permissionChecker(r); permChecker != nil {
		result = append(result, permChecker)
	}
	if s.Middlewares != nil && len(s.Middlewares) > 0 {
		result = append(result, s.Middlewares...)
	}
	return result
}

func setFieldsToEmpty(target, empty interface{}) {
	targetValue := reflect.ValueOf(target).Elem()
	emptyValue := reflect.ValueOf(empty).Elem()

	for i := 0; i < targetValue.NumField(); i++ {
		targetField := targetValue.Field(i)
		emptyField := emptyValue.Field(i)
		name := emptyField.Elem().Type().Name()

		isNested := name == "UserRouteSetting" || name == "RoleRouteSetting" || name == "PermissionRouteSetting"
		if targetField.Kind() == reflect.Struct || isNested {
			if isNested {
				if targetField.IsZero() {
					targetField.Set(emptyField)
				} else {
					setFieldsToEmpty(targetField.Interface(), emptyField.Interface())
				}
				continue
			}
			// If the field is a struct, recursively call the function
			setFieldsToEmpty(targetField.Addr().Interface(), emptyField.Addr().Interface())
		} else if targetField.IsZero() {
			// If the field is zero, set it to the value from the 'emptySetting' struct
			targetField.Set(emptyField)
		}
		if setting, ok := targetField.Interface().(*Setting); ok {
			if setting.ok() {

			}
		}
	}
}

func (r *RouteSetting) validate() {
	setFieldsToEmpty(r, emptySetting)
}

var emptySetting = &RouteSetting{
	Root: &Setting{
		Middlewares: nil,
		Name:        "",
		Permission:  "",
	},
	ListActionsResources: &Setting{
		Middlewares: nil,
		Name:        "",
		Permission:  "",
	},
	Role: &RoleRouteSetting{
		Root: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		ListRolesForUser: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		HasPermission: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		AddInheritance: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		DelInheritance: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		DelRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
	},
	User: &UserRouteSetting{
		Root: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		ListUsersForRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		IsUserIn: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		HasPermission: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		AddRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		DelRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		BulkAddRoles: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		BulkDelRoles: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		DelUser: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
	},
	Permission: &PermissionRouteSetting{
		Root: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		ListPermissionsForUser: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		AddPermissionForUser: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		AddPermissionForRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		DelPermissionForUser: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		DelPermissionForRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		BulkAddPermissionsForUser: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		BulkAddPermissionsForRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		BulkDelPermissionsForUser: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
		BulkDelPermissionsForRole: &Setting{
			Middlewares: nil,
			Name:        "",
			Permission:  "",
		},
	},
	ListCurrentUserPermissions: &Setting{
		Middlewares: nil,
		Name:        "",
		Permission:  "",
	},
}
