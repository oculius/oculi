package rbac

import (
	authz2 "github.com/oculius/oculi/v2/auth/authz"
)

type (
	userReq struct {
		User string `json:"user"`
	}

	roleReq struct {
		Role string `json:"role"`
	}

	twinRoleReq struct {
		ParentRole string `json:"parent_role"`
		ChildRole  string `json:"child_role"`
	}

	userRoleReq struct {
		User string `json:"user"`
		Role string `json:"role"`
	}

	userRolesReq struct {
		User  string   `json:"user"`
		Roles []string `json:"roles"`
	}

	roleBulkPermsReq struct {
		Role        string             `json:"role"`
		Permissions authz2.Permissions `json:"permissions"`
	}

	userBulkPermsReq struct {
		User        string             `json:"user"`
		Permissions authz2.Permissions `json:"permissions"`
	}

	roleSinglePermReq struct {
		Role       string            `json:"role"`
		Permission authz2.Permission `json:"permission"`
	}

	userSinglePermReq struct {
		User       string            `json:"user"`
		Permission authz2.Permission `json:"permission"`
	}
)

func (t twinRoleReq) value() (string, string) {
	return t.ParentRole, t.ChildRole
}

func (u userRolesReq) value() (string, authz2.Roles) {
	return u.User, u.Roles
}

func (r roleReq) value() string {
	return r.Role
}

func (u userReq) value() string {
	return u.User
}

func (u userRoleReq) value() (string, string) {
	return u.User, u.Role
}

func (u userSinglePermReq) value() (string, string, string) {
	return u.User, u.Permission.Resource, u.Permission.Action
}

func (r roleSinglePermReq) value() (string, string, string) {
	return r.Role, r.Permission.Resource, r.Permission.Action
}

func (r roleBulkPermsReq) value() (string, authz2.Permissions) {
	return r.Role, r.Permissions
}

func (u userBulkPermsReq) value() (string, authz2.Permissions) {
	return u.User, u.Permissions
}

var (
	_ doubleDataReq[string, authz2.Permissions] = userBulkPermsReq{}
	_ doubleDataReq[string, authz2.Permissions] = roleBulkPermsReq{}
	_ tripleDataReq                             = roleSinglePermReq{}
	_ tripleDataReq                             = userSinglePermReq{}
	_ doubleDataReq[string, authz2.Roles]       = userRolesReq{}
	_ doubleDataReq[string, string]             = userRoleReq{}
	_ singleDataReq                             = userReq{}
	_ singleDataReq                             = roleReq{}
	_ doubleDataReq[string, string]             = twinRoleReq{}
)
