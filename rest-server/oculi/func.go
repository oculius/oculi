package oculi

import "github.com/labstack/echo/v4"

type (
	HandlerFunc    func(ctx Context) error
	MiddlewareFunc func(next HandlerFunc) HandlerFunc
)

// H stands for Handler Function
func H(handlerFunc HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := C(c)
		return handlerFunc(ctx)
	}
}

// C stands for Context
func C(c echo.Context) Context {
	ctx, ok := c.(Context)
	if !ok {
		panic("oculi context is not found")
	}
	return ctx
}

func echoToOculi(fn echo.HandlerFunc) HandlerFunc {
	return func(ctx Context) error {
		return fn(ctx)
	}
}

// M stands for Middleware
func M(mwFn MiddlewareFunc) echo.MiddlewareFunc {
	return func(nextEcho echo.HandlerFunc) echo.HandlerFunc {
		nextOculi := echoToOculi(nextEcho)
		fnResultOculi := mwFn(nextOculi)
		fnResultEcho := H(fnResultOculi)
		return fnResultEcho
	}
}
