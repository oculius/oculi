package oculi

import (
	"github.com/labstack/echo/v4"
	"io/fs"
)

type (
	Route struct {
		*echo.Route
	}

	group struct {
		*echo.Group
	}

	RouteGroup interface {
		Use(middleware ...MiddlewareFunc)
		CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		Any(path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
		Match(methods []string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
		RouteGroup(prefix string, middleware ...MiddlewareFunc) RouteGroup
		File(path string, file string)
		RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		Add(method string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) Route
		Static(pathPrefix string, fsRoot string)
		StaticFS(pathPrefix string, filesystem fs.FS)
		FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route
	}
)

func NewRoute(r *echo.Route) Route {
	return Route{r}
}
func NewRoutes(r []*echo.Route) []Route {
	if r == nil {
		return nil
	}
	N := len(r)
	result := make([]Route, N)
	for i, each := range r {
		result[i] = NewRoute(each)
	}
	return result
}

func NewRouteGroup(g *echo.Group) RouteGroup {
	return group{g}
}

func (g group) Use(m ...MiddlewareFunc) {
	g.Group.Use(mote(m)...)
}

type generalFunc func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route

func (g group) generalTransform(path string, h HandlerFunc, m []MiddlewareFunc, fn generalFunc) Route {
	return NewRoute(fn(path, H(h), mote(m)...))
}

func (g group) CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.CONNECT)
}

func (g group) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.DELETE)
}

func (g group) GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.GET)
}

func (g group) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.HEAD)
}

func (g group) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.OPTIONS)
}

func (g group) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.PATCH)
}

func (g group) POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.POST)
}

func (g group) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.PUT)
}

func (g group) TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.TRACE)
}

func (g group) Any(path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(g.Group.Any(path, H(h), mote(m)...))
}

func (g group) Match(methods []string, path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(g.Group.Match(methods, path, H(h), mote(m)...))
}

func (g group) RouteGroup(prefix string, middleware ...MiddlewareFunc) RouteGroup {
	return NewRouteGroup(g.Group.Group(prefix, mote(middleware)...))
}

func (g group) RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.RouteNotFound)
}

func (g group) Add(method string, path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return NewRoute(g.Group.Add(method, path, H(h), mote(m)...))
}

func (g group) FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route {
	return NewRoute(g.Group.FileFS(path, file, filesystem, mote(m)...))
}
