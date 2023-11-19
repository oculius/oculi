package oculi

import "github.com/labstack/echo/v4"

// Converter for MiddlewareFunc to echo.MiddlewareFunc
func ToEchoMiddleware(mwFn MiddlewareFunc) echo.MiddlewareFunc {
	return func(nextEcho echo.HandlerFunc) echo.HandlerFunc {
		nextOculi := FromEchoHandler(nextEcho)
		fnResultOculi := mwFn(nextOculi)
		return ToEchoHandler(fnResultOculi)
	}
}

func FromEchoMiddleware(mwFn echo.MiddlewareFunc) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		nextEcho := ToEchoHandler(next)
		fnResultEcho := mwFn(nextEcho)
		fnResultOculi := FromEchoHandler(fnResultEcho)
		return fnResultOculi
	}
}

func bulkToEchoMiddleware(middleware []MiddlewareFunc) []echo.MiddlewareFunc {
	N := len(middleware)
	if N == 0 {
		return []echo.MiddlewareFunc{}
	}
	result := make([]echo.MiddlewareFunc, N)
	for i, each := range middleware {
		result[i] = ToEchoMiddleware(each)
	}
	return result
}
