package oculi

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/http-error"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest/oculi/token"
)

type (
	oculiCtx struct {
		echo.Context
		ctx context.Context
	}
)

// Converter for echo.Context to Context
func FromEchoContext(c echo.Context) Context {
	ctx, ok := c.(Context)
	if !ok {
		panic("oculi context is not found")
	}
	return ctx
}

func (c *oculiCtx) RequestContext() context.Context {
	return c.ctx
}

func (c *oculiCtx) Send(httpResponse response.Convertible) error {
	return c.JSON(httpResponse.ResponseCode(), response.New(httpResponse))
}

func (c *oculiCtx) SendPretty(httpResponse response.Convertible) error {
	return c.JSONPretty(httpResponse.ResponseCode(), response.New(httpResponse), "\t")
}

func (c *oculiCtx) IsDevelopment() bool {
	return c.Get("development") != nil
}

func (c *oculiCtx) AutoSend(httpResponse response.Convertible) error {
	if c.IsDevelopment() {
		return c.SendPretty(httpResponse)
	}
	return c.SendPretty(httpResponse)
}

func NewFromEchoContext(ctx echo.Context) Context {
	return &oculiCtx{
		Context: ctx,
		ctx:     ctx.Request().Context(),
	}
}

func (c *oculiCtx) Lookup(tokens ...token.Token) (map[string]token.Token, httperror.HttpError) {
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

func (c *oculiCtx) BindValidate(obj interface{}) httperror.HttpError {
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
