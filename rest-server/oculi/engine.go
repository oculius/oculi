package oculi

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"io/fs"
	"net"
	"net/http"
)

type (
	Engine struct {
		*echo.Echo
	}

	IEngine interface {
		NewContext(r *http.Request, w http.ResponseWriter) Context
		Router() *echo.Router
		Routers() map[string]*echo.Router
		DefaultHTTPErrorHandler(err error, c echo.Context)
		Use(middleware ...MiddlewareFunc)
		UseEchoMiddleware(middleware ...echo.MiddlewareFunc)
		CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		Any(path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
		Match(methods []string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
		File(path string, file string, m ...MiddlewareFunc) Route
		Add(method string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) Route
		Host(name string, m ...MiddlewareFunc) RouteGroup
		Group(prefix string, m ...MiddlewareFunc) RouteGroup
		RouteGroup(prefix string, m ...MiddlewareFunc) RouteGroup
		URI(handler HandlerFunc, params ...interface{}) string
		URL(h HandlerFunc, params ...interface{}) string
		Reverse(name string, params ...interface{}) string
		Routes() []Route
		AcquireContext() echo.Context
		ReleaseContext(c echo.Context)
		ServeHTTP(w http.ResponseWriter, r *http.Request)
		Start(address string) error
		StartTLS(address string, certFile interface{}, keyFile interface{}) (err error)
		StartAutoTLS(address string) error
		StartServer(s *http.Server) (err error)
		ListenerAddr() net.Addr
		TLSListenerAddr() net.Addr
		StartH2CServer(address string, h2s *http2.Server) error
		Close() error
		Shutdown(ctx context.Context) error
		Static(pathPrefix string, fsRoot string) Route
		StaticFS(pathPrefix string, filesystem fs.FS) Route
		FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route
	}
)

func New() Engine {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			oculiCtx, ok := c.(Context)
			if !ok {
				oculiCtx = NewContext(c)
			}
			return next(oculiCtx)
		}
	})
	return Engine{e}
}

var _ IEngine = &Engine{}

func (e Engine) NewContext(r *http.Request, w http.ResponseWriter) Context {
	return C(e.Echo.NewContext(r, w))
}

func (e Engine) Use(middleware ...MiddlewareFunc) {
	e.Echo.Use(mote(middleware)...)
}

func (e Engine) UseEchoMiddleware(middleware ...echo.MiddlewareFunc) {
	e.Echo.Use(middleware...)
}

func (e Engine) generalTransform(path string, h HandlerFunc, m []MiddlewareFunc, fn generalFunc) Route {
	return NewRoute(fn(path, H(h), mote(m)...), h)
}

func (e Engine) CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.CONNECT)
}

func (e Engine) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.DELETE)
}

func (e Engine) GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.GET)
}

func (e Engine) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.HEAD)
}

func (e Engine) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.OPTIONS)
}

func (e Engine) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.PATCH)
}

func (e Engine) POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.POST)
}

func (e Engine) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.PUT)
}

func (e Engine) TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.TRACE)
}

func (e Engine) RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return e.generalTransform(path, h, m, e.Echo.RouteNotFound)
}

func (e Engine) Any(path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(e.Echo.Any(path, H(h), mote(m)...))
}

func (e Engine) Match(methods []string, path string, h HandlerFunc, m ...MiddlewareFunc) []Route {
	return NewRoutes(e.Echo.Match(methods, path, H(h), mote(m)...))
}

func (e Engine) File(path string, file string, m ...MiddlewareFunc) Route {
	return NewRoute(e.Echo.File(path, file, mote(m)...), e.File)
}

func (e Engine) Add(method string, path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	return NewRoute(e.Echo.Add(method, path, H(h), mote(m)...), h)
}

func (e Engine) Host(name string, m ...MiddlewareFunc) RouteGroup {
	return NewRouteGroup(e.Echo.Host(name, mote(m)...))
}

func (e Engine) Group(prefix string, m ...MiddlewareFunc) RouteGroup {
	return NewRouteGroup(e.Echo.Group(prefix, mote(m)...))
}

func (e Engine) RouteGroup(prefix string, m ...MiddlewareFunc) RouteGroup {
	return NewRouteGroup(e.Echo.Group(prefix, mote(m)...))
}

func (e Engine) URI(h HandlerFunc, params ...interface{}) string {
	return e.Echo.URI(H(h), params...)
}

func (e Engine) URL(h HandlerFunc, params ...interface{}) string {
	return e.Echo.URL(H(h), params...)
}

func (e Engine) Routes() []Route {
	return NewRoutes(e.Echo.Routes())
}

func (e Engine) Static(pathPrefix string, fsRoot string) Route {
	return NewRoute(e.Echo.Static(pathPrefix, fsRoot), e.Static)
}

func (e Engine) StaticFS(pathPrefix string, filesystem fs.FS) Route {
	return NewRoute(e.Echo.StaticFS(pathPrefix, filesystem), e.StaticFS)
}

func (e Engine) FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route {
	return NewRoute(e.Echo.FileFS(path, file, filesystem, mote(m)...), e.FileFS)
}
