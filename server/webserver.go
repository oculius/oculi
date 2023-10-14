package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/oculius/oculi/v2/common/error-extension"
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/server/oculi"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	webServer struct {
		core           Core
		resource       IResource
		config         Config
		useDefaultGZip bool
		signal         chan os.Signal

		afterRun   []HookFunction
		beforeRun  []HookFunction
		beforeExit []HookFunction
		afterExit  []HookFunction
	}
)

func New(core Core, resource IResource, config Config) (Server, error) {
	if core == nil {
		return nil, errors.New("Core is nil")
	}

	if resource == nil {
		return nil, errors.New("Resource is nil")
	}

	if config == nil {
		return nil, errors.New("Config is nil")
	}
	return &webServer{
		core:           core,
		resource:       resource,
		config:         config,
		useDefaultGZip: true,
		signal:         make(chan os.Signal, 3),
	}, nil
}

func (w *webServer) requestPrinter(next oculi.HandlerFunc) oculi.HandlerFunc {
	return func(c oculi.Context) error {
		start := time.Now()
		err := next(c)
		_, errformat := fmt.Fprint(w.resource.Logger().Output(), printerInstance.fmtRequest(c, start, err))
		if errformat != nil {
			w.resource.Logger().OError(
				logs.NewInfo("request formatting error",
					logs.KVs("error", err.Error()),
				),
			)
		}
		return err
	}
}

func (w *webServer) Signal(signal os.Signal) {
	if signal != nil && w.signal != nil {
		w.signal <- signal
	}
}

func (w *webServer) Run() error {
	if err := w.start(); err != nil {
		return err
	}

	signal.Notify(w.signal, syscall.SIGINT, syscall.SIGTERM)
	w.resource.Engine().Echo.HideBanner = true
	if w.config.IsDevelopmentMode() {
		w.resource.Engine().Use(w.requestPrinter)
		w.resource.Engine().Use(func(next oculi.HandlerFunc) oculi.HandlerFunc {
			return func(ctx oculi.Context) error {
				ctx.Set("development", true)
				return next(ctx)
			}
		})
		printerInstance.printRoutes(w.resource.Engine().Echo)
	}
	if err := w.apply(w.beforeRun); err != nil {
		return err
	}

	go w.serve()

	if err := w.apply(w.afterRun); err != nil {
		return err
	}

	w.resource.Logger().Infof("http server started, name: %s", w.resource.ServiceName())
	<-w.signal

	if err := w.apply(w.beforeExit); err != nil {
		return err
	}
	w.stop()
	if err := w.apply(w.afterExit); err != nil {
		return err
	}
	return nil
}

func (w *webServer) start() error {
	echoEngine := w.resource.Engine()
	echoEngine.UseEchoMiddleware(middleware.Recover())
	// TODO Server/Validator
	// echoEngine.Validator = w.boilerplate.Validator()
	echoEngine.Logger = w.resource.Logger()
	echoEngine.Logger.SetLevel(log.INFO)

	if w.useDefaultGZip {
		echoEngine.UseEchoMiddleware(middleware.Gzip())
	}

	echoEngine.UseEchoMiddleware(middleware.RemoveTrailingSlash())

	echo.NotFoundHandler = func(c echo.Context) error {
		err := ErrNotFound(nil, nil)
		return c.JSON(err.ResponseCode(), response.New(err))
	}
	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		err := ErrMethodNotAllowed(nil, nil)
		return c.JSON(err.ResponseCode(), response.New(err))
	}
	echoEngine.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			w.resource.Logger().OError(
				logs.NewInfo("http error handler: response already committed, got error",
					logs.KV("error", err.Error()),
					logs.KV("error_obj", fmt.Sprintf("%+v", err)),
				),
			)
			return
		}

		genericError, ok := err.(errext.Error)
		if !ok {
			httpError, ok2 := err.(*echo.HTTPError)
			if ok2 {
				genericError = errext.New("unknown http error", httpError.Code)(httpError.Internal, httpError.Message)
			} else {
				genericError = errext.New("unknown error", http.StatusInternalServerError)(err, nil)
			}
		}

		var logErr error

		if c.Request().Method == http.MethodHead {
			logErr = c.NoContent(genericError.ResponseCode())
		} else {
			logErr = c.JSON(genericError.ResponseCode(), response.New(genericError))
		}

		if logErr != nil {
			w.resource.Logger().OError(
				logs.NewInfo("http error handler: while committing error response, got error",
					logs.KV("error", err.Error()),
					logs.KV("error_obj", fmt.Sprintf("%+v", err)),
				),
			)
		}
	}

	if err := w.core.Init(echoEngine); err != nil {
		w.resource.Logger().Error("error on init http server")
		return err
	}

	if w.core.Health() != nil {
		echoEngine.GET("/health", w.core.Health())
	}
	return nil
}

func (w *webServer) serve() {
	if err := w.resource.Engine().Start(fmt.Sprintf(":%d", w.resource.ServerPort())); err != nil {
		w.resource.Logger().Infof("http server stopped, %s", err.Error())
		w.signal <- syscall.SIGINT
	} else {
		w.resource.Logger().Info("starting apps")
	}
}

func (w *webServer) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.ServerGracefullyDuration())
	defer func() {
		close(w.signal)
		cancel()
	}()

	if err := w.resource.Engine().Shutdown(ctx); err != nil {
		w.resource.Logger().Errorf("failed to shutdown http server %s", err)
	}
	w.resource.Logger().Info("closing resource")
	if err := w.resource.Close(); err != nil {
		w.resource.Logger().Errorf("failed to close resource %s", err)
	}
}
