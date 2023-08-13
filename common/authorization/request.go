package authorization

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
		Role        string      `json:"role"`
		Permissions Permissions `json:"permissions"`
	}

	userBulkPermsReq struct {
		User        string      `json:"user"`
		Permissions Permissions `json:"permissions"`
	}

	roleSinglePermReq struct {
		Role       string     `json:"role"`
		Permission Permission `json:"permission"`
	}

	userSinglePermReq struct {
		User       string     `json:"user"`
		Permission Permission `json:"permission"`
	}

	tripleDataReq interface {
		value() (target string, resource string, action string)
	}

	doubleDataReq[T any, V any] interface {
		value() (T, V)
	}

	singleDataReq interface {
		value() string
	}
)

func (t twinRoleReq) value() (string, string) {
	return t.ParentRole, t.ChildRole
}

func (u userRolesReq) value() (string, Roles) {
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

func (r roleBulkPermsReq) value() (string, Permissions) {
	return r.Role, r.Permissions
}

func (u userBulkPermsReq) value() (string, Permissions) {
	return u.User, u.Permissions
}

var (
	_ doubleDataReq[string, Permissions] = userBulkPermsReq{}
	_ doubleDataReq[string, Permissions] = roleBulkPermsReq{}
	_ tripleDataReq                      = roleSinglePermReq{}
	_ tripleDataReq                      = userSinglePermReq{}
	_ doubleDataReq[string, Roles]       = userRolesReq{}
	_ doubleDataReq[string, string]      = userRoleReq{}
	_ singleDataReq                      = userReq{}
	_ singleDataReq                      = roleReq{}
	_ doubleDataReq[string, string]      = twinRoleReq{}
)
