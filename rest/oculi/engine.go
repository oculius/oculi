package oculi

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Engine struct {
		*echo.Echo
	}
)

func oculiSpawner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, ok := c.(Context)
		if !ok {
			ctx = NewFromEchoContext(c)
		}
		return next(ctx)
	}
}

func New() *Engine {
	e := echo.New()
	e.Use(oculiSpawner)
	return &Engine{e}
}

//var _ Engine = &engine{}

func (e *Engine) NewContext(r *http.Request, w http.ResponseWriter) Context {
	return FromEchoContext(e.Echo.NewContext(r, w))
}

func (e *Engine) Use(middleware ...MiddlewareFunc) {
	e.Echo.Use(bulkToEchoMiddleware(middleware)...)
}

func (e *Engine) UseEchoMiddleware(middleware ...echo.MiddlewareFunc) {
	e.Echo.Use(middleware...)
}

func (e *Engine) generalTransform(path string, h HandlerFunc, m []MiddlewareFunc, fn generalFunc) Route {
	return FromEchoRoute(fn(path, ToEchoHandler(h), bulkToEchoMiddleware(m)...), h)
}

func (e *Engine) CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.CONNECT)
}

func (e *Engine) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.DELETE)
}

func (e *Engine) GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.GET)
}

func (e *Engine) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.HEAD)
}

func (e *Engine) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.OPTIONS)
}

func (e *Engine) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.PATCH)
}

func (e *Engine) POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.POST)
}

func (e *Engine) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.PUT)
}

func (e *Engine) TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.TRACE)
}

func (e *Engine) RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.RouteNotFound)
}

func (e *Engine) Any(path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(e.Echo.Any(path, ToEchoHandler(h), bulkToEchoMiddleware(m)...))
}

func (e *Engine) Match(methods []string, path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(e.Echo.Match(methods, path, ToEchoHandler(h), bulkToEchoMiddleware(m)...))
}

func (e *Engine) File(path string, file string, m ...MiddlewareFunc) Route {
	return FromEchoRoute(e.Echo.File(path, file, bulkToEchoMiddleware(m)...), e.File)
}

func (e *Engine) Add(method string, path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return FromEchoRoute(e.Echo.Add(method, path, ToEchoHandler(h), bulkToEchoMiddleware(m)...), h)
}

func (e *Engine) Host(name string, m ...MiddlewareFunc) RouteGroup {
	return FromEchoGroup(e.Echo.Host(name, bulkToEchoMiddleware(m)...))
}

func (e *Engine) Group(prefix string, m ...MiddlewareFunc) RouteGroup {
	return FromEchoGroup(e.Echo.Group(prefix, bulkToEchoMiddleware(m)...))
}

func (e *Engine) RouteGroup(prefix string, m ...MiddlewareFunc) RouteGroup {
	return FromEchoGroup(e.Echo.Group(prefix, bulkToEchoMiddleware(m)...))
}

func (e *Engine) URI(h HandlerFunc, params ...interface{}) string {
	return e.Echo.URI(ToEchoHandler(h), params...)
}

func (e *Engine) URL(h HandlerFunc, params ...interface{}) string {
	return e.Echo.URL(ToEchoHandler(h), params...)
}

func (e *Engine) Routes() []Route {
	return NewRoutes(e.Echo.Routes())
}

func (e *Engine) Static(pathPrefix string, fsRoot string) Route {
	return FromEchoRoute(e.Echo.Static(pathPrefix, fsRoot), e.Static)
}

func (e *Engine) StaticFS(pathPrefix string, filesystem fs.FS) Route {
	return FromEchoRoute(e.Echo.StaticFS(pathPrefix, filesystem), e.StaticFS)
}

func (e *Engine) FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route {
	return FromEchoRoute(e.Echo.FileFS(path, file, filesystem, bulkToEchoMiddleware(m)...), e.FileFS)
}
