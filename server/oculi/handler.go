package oculi

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	HandlerFunc    func(ctx Context) error
	MiddlewareFunc func(next HandlerFunc) HandlerFunc
)

// H stands for Handler Function, Converter for HandlerFunc to echo.HandlerFunc
func H(handlerFunc HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := C(c)
		return handlerFunc(ctx)
	}
}

func WrapHandler(handler http.HandlerFunc) HandlerFunc {
	return NewHandlerFunc(echo.WrapHandler(handler))
}

// C stands for Context, Converter for echo.Context to Context
func C(c echo.Context) Context {
	ctx, ok := c.(Context)
	if !ok {
		panic("oculi context is not found")
	}
	return ctx
}

func NewHandlerFunc(fn echo.HandlerFunc) HandlerFunc {
	return func(ctx Context) error {
		return fn(ctx)
	}
}

// M stands for Middleware, Converter for MiddlewareFunc to echo.MiddlewareFunc
func M(mwFn MiddlewareFunc) echo.MiddlewareFunc {
	return func(nextEcho echo.HandlerFunc) echo.HandlerFunc {
		nextOculi := NewHandlerFunc(nextEcho)
		fnResultOculi := mwFn(nextOculi)
		return H(fnResultOculi)
	}
}

func NewMiddlewareFunc(mwFn echo.MiddlewareFunc) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		nextEcho := H(next)
		fnResultEcho := mwFn(nextEcho)
		fnResultOculi := NewHandlerFunc(fnResultEcho)
		return fnResultOculi
	}
}

// mote stands for middleware oculi to echo
func mote(middleware []MiddlewareFunc) []echo.MiddlewareFunc {
	N := len(middleware)
	if N == 0 {
		return []echo.MiddlewareFunc{}
	}
	result := make([]echo.MiddlewareFunc, N)
	for i, each := range middleware {
		result[i] = M(each)
	}
	return result
}
