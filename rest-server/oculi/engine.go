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
		Pre(middleware ...MiddlewareFunc)
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
		RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		Any(path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
		Match(methods []string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
		File(path string, file string, m ...MiddlewareFunc) Route
		Add(method string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) Route
		Host(name string, m ...MiddlewareFunc) RouteGroup
		Group(prefix string, m ...MiddlewareFunc) RouteGroup
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

func New(e *echo.Echo) Engine {
	return Engine{e}
}

var _ IEngine = Engine{}

func (e Engine) NewContext(r *http.Request, w http.ResponseWriter) Context {
	return C(e.Echo.NewContext(r, w))
}

func (e Engine) Pre(middleware ...MiddlewareFunc) {
	e.Echo.Pre(mote(middleware)...)
}

func (e Engine) Use(middleware ...MiddlewareFunc) {
	e.Echo.Use(mote(middleware)...)
}

func (e Engine) CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) Any(path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) Match(methods []string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) File(path string, file string, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) Add(method string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) Host(name string, m ...MiddlewareFunc) RouteGroup {
	//TODO implement me
	panic("implement me")
}

func (e Engine) Group(prefix string, m ...MiddlewareFunc) RouteGroup {
	//TODO implement me
	panic("implement me")
}

func (e Engine) URI(handler HandlerFunc, params ...interface{}) string {
	//TODO implement me
	panic("implement me")
}

func (e Engine) URL(h HandlerFunc, params ...interface{}) string {
	//TODO implement me
	panic("implement me")
}

func (e Engine) Routes() []Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) Static(pathPrefix string, fsRoot string) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) StaticFS(pathPrefix string, filesystem fs.FS) Route {
	//TODO implement me
	panic("implement me")
}

func (e Engine) FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route {
	//TODO implement me
	panic("implement me")
}
