package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/oculius/oculi/v2/application/logs"
	"github.com/oculius/oculi/v2/common/http-error"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest/oculi"
	"github.com/pkg/errors"
)

type (
	webServer struct {
		core           Core
		option         Option
		useDefaultGZip bool
		signal         chan os.Signal
		engine         *oculi.Engine
		logger         logs.Logger

		afterRun   []HookFunction
		beforeRun  []HookFunction
		beforeExit []HookFunction
		afterExit  []HookFunction
	}
)

func New[X Core](core X, opt Option, engine *oculi.Engine, logger logs.Logger) (Server, error) {
	if any(core) == nil {
		return nil, errors.New("Core is nil")
	}

	return &webServer{
		core:           core,
		option:         opt,
		useDefaultGZip: true,
		engine:         engine,
		logger:         logger,
		signal:         make(chan os.Signal, 3),
	}, nil
}

func (w *webServer) requestPrinter(next oculi.HandlerFunc) oculi.HandlerFunc {
	return func(c oculi.Context) error {
		start := time.Now()
		err := next(c)
		_, errformat := fmt.Fprint(w.logger.Output(), printerInstance.fmtRequest(c, start, err))
		if errformat != nil {
			w.logger.OError(
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
	w.engine.Echo.HideBanner = true
	if w.option.IsDevMode {
		w.engine.Use(w.requestPrinter)
		w.engine.Use(func(next oculi.HandlerFunc) oculi.HandlerFunc {
			return func(ctx oculi.Context) error {
				ctx.Set("development", true)
				return next(ctx)
			}
		})
		printerInstance.printRoutes(w.engine.Echo)
	}
	if err := w.apply(w.beforeRun); err != nil {
		return err
	}

	go w.serve()

	if err := w.apply(w.afterRun); err != nil {
		return err
	}

	w.logger.Infof("http server started, name: %s", w.option.ServiceName)
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
	echoEngine := w.engine
	echoEngine.UseEchoMiddleware(middleware.Recover())
	// TODO Server/Validator
	// echoEngine.Validator = w.boilerplate.Validator()
	echoEngine.Logger = w.logger
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
			w.logger.OError(
				logs.NewInfo("http error handler: response already committed, got error",
					logs.KV("error", err.Error()),
					logs.KV("error_obj", fmt.Sprintf("%+v", err)),
				),
			)
			return
		}

		genericError, ok := err.(httperror.HttpError)
		if !ok {
			httpError, ok2 := err.(*echo.HTTPError)
			if ok2 {
				genericError = httperror.New("unknown http error", httpError.Code)(httpError.Internal, httpError.Message)
			} else {
				genericError = httperror.New("unknown error", http.StatusInternalServerError)(err, nil)
			}
		}

		var logErr error

		if c.Request().Method == http.MethodHead {
			logErr = c.NoContent(genericError.ResponseCode())
		} else {
			logErr = c.JSON(genericError.ResponseCode(), response.New(genericError))
		}

		if logErr != nil {
			w.logger.OError(
				logs.NewInfo("http error handler: while committing error response, got error",
					logs.KV("error", err.Error()),
					logs.KV("error_obj", fmt.Sprintf("%+v", err)),
				),
			)
		}
	}

	if err := w.core.Init(echoEngine); err != nil {
		w.logger.Error("error on init http server")
		return err
	}

	if w.core.Healthcheck != nil {
		echoEngine.GET("/health", w.core.Healthcheck)
	}
	return nil
}

func (w *webServer) serve() {
	if err := w.engine.Start(fmt.Sprintf(":%d", w.option.ServicePort)); err != nil {
		w.logger.Infof("http server stopped, %s", err.Error())
		w.signal <- syscall.SIGINT
	} else {
		w.logger.Info("starting apps")
	}
}

func (w *webServer) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), w.option.ShutdownGracefulDuration)
	defer func() {
		close(w.signal)
		cancel()
	}()

	if err := w.engine.Shutdown(ctx); err != nil {
		w.logger.Errorf("failed to shutdown http server: %s", err)
	}
	w.logger.Info("closing resource")
	if err := w.engine.Close(); err != nil {
		w.logger.Errorf("failed to close http server: %s", err)
	}
}
