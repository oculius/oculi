package authz

import (
	"github.com/casbin/casbin/v2"
	"github.com/oculius/oculi/v2/rest/oculi"
)

type (
	Enforcer   casbin.IEnforcer
	TxFunction func(enf Enforcer) error

	CasbinService interface {
		Enforcer() Enforcer
		Transaction(fn TxFunction) error
	}

	Roles []string
	Users []string

	UserRetrieverREST func(ctx oculi.Context) string
)

var (
	PublicIdentifier = "guest"
)
