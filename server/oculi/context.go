package oculi

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/http-error"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/server/oculi/token"
)

type (
	oculiContext struct {
		echo.Context
		ctx context.Context
	}

	Context interface {
		echo.Context
		BindValidate(interface{}) httperror.HttpError
		Lookup(...token.Token) (map[string]token.Token, httperror.HttpError)
		Send(response.Convertible) error
		SendPretty(response.Convertible) error
		IsDevelopment() bool
		AutoSend(response.Convertible) error
		RequestContext() context.Context
	}
)

func (c *oculiContext) RequestContext() context.Context {
	return c.ctx
}

func (c *oculiContext) Send(httpResponse response.Convertible) error {
	return c.JSON(httpResponse.ResponseCode(), response.New(httpResponse))
}

func (c *oculiContext) SendPretty(httpResponse response.Convertible) error {
	return c.JSONPretty(httpResponse.ResponseCode(), response.New(httpResponse), "\t")
}

func (c *oculiContext) IsDevelopment() bool {
	return c.Get("development") != nil
}

func (c *oculiContext) AutoSend(httpResponse response.Convertible) error {
	if c.IsDevelopment() {
		return c.SendPretty(httpResponse)
	}
	return c.SendPretty(httpResponse)
}

func NewContext(echoCtx echo.Context) Context {
	return &oculiContext{
		Context: echoCtx,
		ctx:     echoCtx.Request().Context(),
	}
}

func (c *oculiContext) Lookup(tokens ...token.Token) (map[string]token.Token, httperror.HttpError) {
	N := len(tokens)
	if N == 0 {
		return nil, nil
	}
	result := make(map[string]token.Token, N)
	for _, t := range tokens {
		err := t.Apply(c.Context)
		if err != nil {
			return nil, err
		}
		result[t.Key()] = t
	}
	return result, nil
}

func (c *oculiContext) BindValidate(obj interface{}) httperror.HttpError {
	if err := c.Bind(obj); err != nil {
		return ErrDataBinding(err, nil)
	}

	if err := c.Validate(obj); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			return ErrDataValidation(err, nil)
		}
		return httperror.NewValidationError(err, Translator)
	}
	return nil
}
