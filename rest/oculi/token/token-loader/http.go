package tl

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/http-error"
	"github.com/pkg/errors"
)

type (
	query        struct{}
	urlparameter struct{}
	header       struct{}
	cookie       struct{}
)

func (c cookie) Load(ctx echo.Context, t Token) (string, httperror.HttpError) {
	if ctx.Request() == nil {
		return "", ErrRequestNotFound(nil, t.Metadata())
	}
	cookie, err := ctx.Cookie(t.Key())
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		return "", nil
	} else if err != nil && !errors.Is(err, http.ErrNoCookie) {
		return "", ErrCookie(err, t.Metadata())
	}
	return cookie.Value, nil
}

func (h header) Load(ctx echo.Context, t Token) (string, httperror.HttpError) {
	if ctx.Request() == nil {
		return "", ErrRequestNotFound(nil, t.Metadata())
	}
	return ctx.Request().Header.Get(t.Key()), nil
}

func (u urlparameter) Load(ctx echo.Context, t Token) (string, httperror.HttpError) {
	return ctx.Param(t.Key()), nil
}

func (q query) Load(ctx echo.Context, t Token) (string, httperror.HttpError) {
	return ctx.QueryParam(t.Key()), nil
}

func Query() Loader[string] {
	return query{}
}

func URLParameter() Loader[string] {
	return urlparameter{}
}

func Header() Loader[string] {
	return header{}
}

func Cookie() Loader[string] {
	return cookie{}
}
