package oculi

import (
	"github.com/labstack/echo/v4"
)

type (
	Route struct {
		*echo.Route
	}

	generalFunc func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route
)

func FromEchoRoute(r *echo.Route, handler any) Route {
	if handler != nil {
		r.Name = handlerName(handler)
	}
	return Route{r}
}
func NewRoutes(r []*echo.Route) []Route {
	if r == nil {
		return nil
	}
	N := len(r)
	result := make([]Route, N)
	for i, each := range r {
		result[i] = FromEchoRoute(each, nil)
	}
	return result
}
