package authz

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func NewCasbinEnforcer(authModel string, db *gorm.DB) Enforcer {
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
