package oculi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Converter for HandlerFunc to echo.HandlerFunc
func ToEchoHandler(handlerFunc HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := FromEchoContext(c)
		return handlerFunc(ctx)
	}
}

func FromHttpHandler(handler http.HandlerFunc) HandlerFunc {
	return FromEchoHandler(echo.WrapHandler(handler))
}

func FromEchoHandler(fn echo.HandlerFunc) HandlerFunc {
	return func(ctx Context) error {
		return fn(ctx)
	}
}
