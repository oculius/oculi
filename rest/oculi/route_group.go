package oculi

import (
	"io/fs"

	"github.com/labstack/echo/v4"
)

type (
	routeGroup struct {
		*echo.Group
	}
)

func FromEchoGroup(g *echo.Group) RouteGroup {
	return &routeGroup{g}
}

func (g *routeGroup) Use(m ...MiddlewareFunc) {
	g.Group.Use(bulkToEchoMiddleware(m)...)
}

func (g *routeGroup) generalTransform(path string, h HandlerFunc, m []MiddlewareFunc, fn generalFunc) Route {
	return FromEchoRoute(fn(path, ToEchoHandler(h), bulkToEchoMiddleware(m)...), h)
}

func (g *routeGroup) CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.CONNECT)
}

func (g *routeGroup) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.DELETE)
}

func (g *routeGroup) GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.GET)
}

func (g *routeGroup) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.HEAD)
}

func (g *routeGroup) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.OPTIONS)
}

func (g *routeGroup) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.PATCH)
}

func (g *routeGroup) POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.POST)
}

func (g *routeGroup) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.PUT)
}

func (g *routeGroup) TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.TRACE)
}

func (g *routeGroup) Any(path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(g.Group.Any(path, ToEchoHandler(h), bulkToEchoMiddleware(m)...))
}

func (g *routeGroup) Match(methods []string, path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(g.Group.Match(methods, path, ToEchoHandler(h), bulkToEchoMiddleware(m)...))
}

func (g *routeGroup) RGroup(prefix string, middleware ...MiddlewareFunc) RouteGroup {
	return FromEchoGroup(g.Group.Group(prefix, bulkToEchoMiddleware(middleware)...))
}

func (g *routeGroup) Bundle(prefix string, bundler func(RouteGroup), middleware ...MiddlewareFunc) RouteGroup {
	rg := FromEchoGroup(g.Group.Group(prefix, bulkToEchoMiddleware(middleware)...))
	bundler(rg)
	return rg
}

func (g *routeGroup) RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return g.generalTransform(path, h, m, g.Group.RouteNotFound)
}

func (g *routeGroup) Add(method string, path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return FromEchoRoute(g.Group.Add(method, path, ToEchoHandler(h), bulkToEchoMiddleware(m)...), h)
}

func (g *routeGroup) FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route {
	return FromEchoRoute(g.Group.FileFS(path, file, filesystem, bulkToEchoMiddleware(m)...), g.FileFS)
}
