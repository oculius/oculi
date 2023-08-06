package oculi

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/error-extension"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server/oculi/token"
)

type (
	oculiContext struct {
		echo.Context
		ctx context.Context
	}

	Context interface {
		echo.Context
		BindValidate(interface{}) errext.Error
		Lookup(...token.Token) (map[string]token.Token, errext.Error)
		Send(response.Convertible) error
		SendPretty(response.Convertible) error
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

func NewContext(echoCtx echo.Context) Context {
	return &oculiContext{
		Context: echoCtx,
		ctx:     echoCtx.Request().Context(),
	}
}

func (c *oculiContext) Lookup(tokens ...token.Token) (map[string]token.Token, errext.Error) {
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

func (c *oculiContext) BindValidate(obj interface{}) errext.Error {
	if err := c.Bind(obj); err != nil {
		return ErrDataBinding(err, nil)
	}

	if err := c.Validate(obj); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			return ErrDataValidation(err, nil)
		}
		return errext.NewValidationError(err, Translator)
	}
	return nil
}
