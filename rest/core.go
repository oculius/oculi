package rest

import (
	"net/http"

	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest/oculi"
)

type (
	defaultCore struct {
		healthcheckHandler oculi.HandlerFunc
		internals          []InternalComponent
		externals          []ExternalComponent
	}
)

func NewCore(healthcheck oculi.HandlerFunc, internals []InternalComponent, externals []ExternalComponent) Core {
	return &defaultCore{healthcheck, internals, externals}
}

func (m *defaultCore) Init(engine *oculi.Engine) error {
	for _, cmp := range m.internals {
		if err := cmp.Init(engine); err != nil {
			return err
		}
	}
	return nil
}

func (m *defaultCore) Healthcheck(ctx oculi.Context) error {
	if m.healthcheckHandler == nil {
		data := map[string]map[string]bool{
			"externals": {},
		}
		anyError := false
		for _, ext := range m.externals {
			err := ext.Ping(ctx.RequestContext())
			data["externals"][ext.Identifier()] = err == nil
			anyError = anyError || err != nil
		}
		if anyError {
			return ctx.AutoSend(
				response.NewResponseWithStatus(
					"NOT OK", data, nil, http.StatusInternalServerError,
				),
			)
		}
		return ctx.AutoSend(response.NewResponse("OK", data, nil))
	}
	return m.healthcheckHandler(ctx)
}
