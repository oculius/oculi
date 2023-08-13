package authorization

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"gorm.io/gorm"
	"strings"
)

type (
	Enforcer   casbin.IEnforcer
	TxFunction func(enf Enforcer) error

	Engine interface {
		rest.Module

		Permission(permission string) oculi.MiddlewareFunc
		ListUsersForRole(ctx oculi.Context) error
		ListRolesForUser(ctx oculi.Context) error
		HasRolesForUser(ctx oculi.Context) error
		AddRolesForUser(ctx oculi.Context) error
		DelRoleForUser(ctx oculi.Context) error
	}
	Service interface {
		Enforcer() Enforcer
		Transaction(fn TxFunction) error
	}

	Permission struct {
		Resource string `json:"resource"`
		Action   string `json:"action"`
	}

	Permissions []Permission
	Roles       []string
	Users       []string

	UserRetrieverREST func(ctx oculi.Context) string
)

const (
	RolePrefix     = "role_"
	ResourcePrefix = "rsrc_"
	UserPrefix     = "user_"
	ActionPrefix   = "actn_"
)

var PublicIdentifier = "guest"

func newEnforcer(authModel string, db *gorm.DB) Enforcer {
	m, err := model.NewModelFromString(authModel)
	if err != nil {
		panic("failed to load authorization model: " + err.Error())
	}
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic("failed to start authorization: " + err.Error())
	}

	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic("failed to instantiate enforcer: " + err.Error())
	}
	return enforcer
}

func (p Permissions) translate(subject string) [][]string {
	if len(p) == 0 {
		return nil
	}
	result := make([][]string, len(p))
	for i, each := range p {
		result[i] = []string{subject, ResourcePrefix + each.Resource, ActionPrefix + each.Action}
	}
	return result
}

func newPermissions(matrix [][]string) Permissions {
	if len(matrix) == 0 {
		return nil
	}
	result := make([]Permission, len(matrix))
	for i, each := range matrix {
		result[i] = Permission{
			Resource: strings.Replace(each[0], ResourcePrefix, "", 1),
			Action:   strings.Replace(each[1], ActionPrefix, "", 1),
		}
	}
	return result
}
