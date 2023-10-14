package authz

import (
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/server/oculi"
	"github.com/oculius/oculi/v2/server/oculi/token"
)

type (
	bulkPermEndPointFunc     func(string, Permissions) (bool, error)
	tripleDataEndPointFunc   func(string, string, string) (bool, error)
	dataCheckerEndPointFunc  func(string, string, string) bool
	twinStringEndPointFunc   func(string, string) (bool, error)
	stringRolesEndPointFunc  func(string, Roles) (bool, error)
	singleStringEndPointFunc func(string) (bool, error)
)

func bulkPermEndpoint[T doubleDataReq[string, Permissions]](fn bulkPermEndPointFunc, ctx oculi.Context) error {
	var body T
	if err := ctx.BindValidate(&body); err != nil {
		return err
	}

	if ok, err := fn(body.value()); err != nil {
		return err
	} else if !ok {
		return ctx.AutoSend(response.NewResponse("failed", nil, nil))
	}

	return ctx.AutoSend(response.NewResponse("success", nil, nil))
}

func tripleDataEndpoint[T tripleDataReq](fn tripleDataEndPointFunc, ctx oculi.Context) error {
	var body T
	if err := ctx.BindValidate(&body); err != nil {
		return err
	}

	if ok, err := fn(body.value()); err != nil {
		return err
	} else if !ok {
		return ctx.AutoSend(response.NewResponse("failed", nil, nil))
	}

	return ctx.AutoSend(response.NewResponse("success", nil, nil))
}

func dataCheckerEndpoint[T tripleDataReq](fn dataCheckerEndPointFunc, ctx oculi.Context) error {
	var body T
	if err := ctx.BindValidate(&body); err != nil {
		return err
	}

	ok := fn(body.value())
	return ctx.AutoSend(response.NewResponse("success", map[string]any{"result": ok}, nil))
}

func twinStringEndpoint[T doubleDataReq[string, string]](fn twinStringEndPointFunc, ctx oculi.Context) error {
	var body T
	if err := ctx.BindValidate(&body); err != nil {
		return err
	}

	if ok, err := fn(body.value()); err != nil {
		return err
	} else if !ok {
		return ctx.AutoSend(response.NewResponse("failed", nil, nil))
	}

	return ctx.AutoSend(response.NewResponse("success", nil, nil))
}

func twinStringCheckerEndpoint[T doubleDataReq[string, string]](fn twinStringEndPointFunc, ctx oculi.Context) error {
	var body T
	if err := ctx.BindValidate(&body); err != nil {
		return err
	}

	ok, err := fn(body.value())
	if err != nil {
		return err
	}

	return ctx.AutoSend(response.NewResponse("success", map[string]any{"result": ok}, nil))
}

func stringRolesEndpoint[T doubleDataReq[string, Roles]](fn stringRolesEndPointFunc, ctx oculi.Context) error {
	var body T
	if err := ctx.BindValidate(&body); err != nil {
		return err
	}

	if ok, err := fn(body.value()); err != nil {
		return err
	} else if !ok {
		return ctx.AutoSend(response.NewResponse("failed", nil, nil))
	}

	return ctx.AutoSend(response.NewResponse("success", nil, nil))
}

func singleStringEndpoint[T singleDataReq](fn singleStringEndPointFunc, ctx oculi.Context) error {
	var body T
	if err := ctx.BindValidate(&body); err != nil {
		return err
	}

	if ok, err := fn(body.value()); err != nil {
		return err
	} else if !ok {
		return ctx.AutoSend(response.NewResponse("failed", nil, nil))
	}

	return ctx.AutoSend(response.NewResponse("success", nil, nil))
}

func getterHelper(ctx oculi.Context, key string, fn func(string) (any, error)) error {
	tokenMap, err := ctx.Lookup(token.T(token.Parameter, key, token.String, true))
	if err != nil {
		return err
	}
	target, err := token.TokenValue[string](tokenMap[key])
	if err != nil {
		return err
	}

	result, errx := fn(target)
	if errx != nil {
		return errx
	}
	return ctx.AutoSend(response.NewResponse("success", result, nil))
}
