package tf

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/error-extension"
	"github.com/pkg/errors"
)

type (
	queryFetcher        struct{}
	urlparameterFetcher struct{}
	headerFetcher       struct{}
	cookieFetcher       struct{}
)

func (c cookieFetcher) Fetch(ctx echo.Context, t Token) (string, errext.HttpError) {
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

func (h headerFetcher) Fetch(ctx echo.Context, t Token) (string, errext.HttpError) {
	if ctx.Request() == nil {
		return "", ErrRequestNotFound(nil, t.Metadata())
	}
	return ctx.Request().Header.Get(t.Key()), nil
}

func (u urlparameterFetcher) Fetch(ctx echo.Context, t Token) (string, errext.HttpError) {
	return ctx.Param(t.Key()), nil
}

func (q queryFetcher) Fetch(ctx echo.Context, t Token) (string, errext.HttpError) {
	return ctx.QueryParam(t.Key()), nil
}

func QueryFetcher() Fetcher[string] {
	return queryFetcher{}
}

func URLParameterFetcher() Fetcher[string] {
	return urlparameterFetcher{}
}

func HeaderFetcher() Fetcher[string] {
	return headerFetcher{}
}

func CookieFetcher() Fetcher[string] {
	return cookieFetcher{}
}

/*
	case Header:
		if ctx.Request() == nil {
			errResult = ErrRequestNotFound(nil, t.Metadata())
			return
		}
		value = ctx.Request().Header.Get(t.key)
	case Cookie:

*/
