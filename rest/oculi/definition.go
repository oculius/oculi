package oculi

import (
	"context"
	"io/fs"

	"github.com/labstack/echo/v4"
	httperror "github.com/oculius/oculi/v2/common/http-error"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest/oculi/token"
)

type (
	Context interface {
		echo.Context
		BindValidate(interface{}) httperror.HttpError
		Lookup(...token.Token) (map[string]token.Token, httperror.HttpError)
		Send(response.Convertible) error
		SendPretty(response.Convertible) error
		IsDevelopment() bool
		AutoSend(response.Convertible) error
		RequestContext() context.Context
	}

	//Engine[Origin any] interface {
	//	NewContext(r *http.Request, w http.ResponseWriter) Context
	//	Origin() Origin
	//	Router() *echo.Router
	//	Routers() map[string]*echo.Router
	//	DefaultHTTPErrorHandler(err error, c echo.Context)
	//	Use(middleware ...MiddlewareFunc)
	//	UseEchoMiddleware(middleware ...echo.MiddlewareFunc)
	//	CONNECT(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	DELETE(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	GET(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	HEAD(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	OPTIONS(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	PATCH(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	POST(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	PUT(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	TRACE(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route
	//	Any(path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
	//	Match(methods []string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) []Route
	//	File(path string, file string, m ...MiddlewareFunc) Route
	//	Add(method string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) Route
	//	Host(name string, m ...MiddlewareFunc) RouteGroup
	//	Group(prefix string, m ...MiddlewareFunc) RouteGroup
	//	RouteGroup(prefix string, m ...MiddlewareFunc) RouteGroup
	//	URI(handler HandlerFunc, params ...interface{}) string
	//	URL(h HandlerFunc, params ...interface{}) string
	//	Reverse(name string, params ...interface{}) string
	//	Routes() []Route
	//	AcquireContext() echo.Context
	//	ReleaseContext(c echo.Context)
	//	ServeHTTP(w http.ResponseWriter, r *http.Request)
	//	Start(address string) error
	//	StartTLS(address string, certFile interface{}, keyFile interface{}) (err error)
	//	StartAutoTLS(address string) error
	//	StartServer(s *http.Server) (err error)
	//	ListenerAddr() net.Addr
	//	TLSListenerAddr() net.Addr
	//	StartH2CServer(address string, h2s *http2.Server) error
	//	Close() error
	//	Shutdown(ctx context.Context) error
	//	Static(pathPrefix string, fsRoot string) Route
	//	StaticFS(pathPrefix string, filesystem fs.FS) Route
	//	FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route
	//}

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
		RGroup(prefix string, middleware ...MiddlewareFunc) RouteGroup
		Bundle(prefix string, arranger func(RouteGroup), middleware ...MiddlewareFunc) RouteGroup
		File(path string, file string)
		RouteNotFound(path string, h HandlerFunc, m ...MiddlewareFunc) Route
		Add(method string, path string, handler HandlerFunc, middleware ...MiddlewareFunc) Route
		Static(pathPrefix string, fsRoot string)
		StaticFS(pathPrefix string, filesystem fs.FS)
		FileFS(path string, file string, filesystem fs.FS, m ...MiddlewareFunc) Route
	}

	HandlerFunc    func(ctx Context) error
	MiddlewareFunc func(next HandlerFunc) HandlerFunc
)
