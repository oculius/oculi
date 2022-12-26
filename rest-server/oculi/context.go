package oculi

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	gerr "github.com/oculius/oculi/v2/common/error"
	"net/http"
)

type (
	oculiContext struct {
		echo.Context
		ctx context.Context
	}

	Context interface {
		echo.Context
		BindValidate(interface{}) gerr.Error
		Lookup(...Token) (map[string]Token, gerr.Error)
	}
)

var (
	FailedBindingHttpStatus               = http.StatusBadRequest
	Translator              ut.Translator = nil
)

func New(echoCtx echo.Context) Context {
	return &oculiContext{
		Context: echoCtx,
		ctx:     echoCtx.Request().Context(),
	}
}

func (c *oculiContext) Lookup(tokens ...Token) (map[string]Token, gerr.Error) {
	N := len(tokens)
	if N == 0 {
		return nil, nil
	}
	result := make(map[string]Token, N)
	for _, t := range tokens {
		err := t.Apply(c.Context)
		if err != nil {
			return nil, err
		}
		result[t.Key()] = t
	}
	return result, nil
}

var (
	ErrDataBinding    = gerr.New("data binding failed", FailedBindingHttpStatus)
	ErrDataValidation = gerr.New("data validation failed", http.StatusInternalServerError)
)

func (c *oculiContext) BindValidate(obj interface{}) gerr.Error {
	if err := c.Bind(obj); err != nil {
		return ErrDataBinding(err, nil)
	}

	if err := c.Validate(obj); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			return ErrDataValidation(err, nil)
		}
		return gerr.NewValidationError(err, Translator)
	}
	return nil
}
