package rest_server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	gerr "github.com/oculius/oculi/v2/common/error"
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func (w *WebServer) DevelopmentMode() {
	w.resource.Echo().Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO Development Logger
			// start := time.Now()
			err := next(c)
			// fmt.Println(formatRequest(c, start))
			return err
		}
	})
}

func (w *WebServer) Run() error {
	if err := w.start(); err != nil {
		return err
	}

	sig := make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	if err := w.apply(w.beforeRun); err != nil {
		return err
	}

	go w.serve(sig)

	if err := w.apply(w.afterRun); err != nil {
		return err
	}
	<-sig

	if err := w.apply(w.beforeExit); err != nil {
		return err
	}
	w.stop()
	if err := w.apply(w.afterExit); err != nil {
		return err
	}
	return nil
}

func (w *WebServer) start() error {
	echoEngine := w.resource.Echo()
	echoEngine.Use(middleware.Recover())
	// TODO Server/Validator
	// echoEngine.Validator = w.resource.Validator()
	echoEngine.Logger = w.resource.Logger()
	echoEngine.Logger.SetLevel(log.INFO)

	if w.useDefaultGZip {
		echoEngine.Use(middleware.Gzip())
	}

	echoEngine.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			oculiCtx, ok := c.(oculi.Context)
			if !ok {
				oculiCtx = oculi.New(c)
			}
			return next(oculiCtx)
		}
	})

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

		genericError, ok := err.(gerr.Error)
		if !ok {
			httpError, ok2 := err.(*echo.HTTPError)
			if ok2 {
				genericError = gerr.New("unknown http error", httpError.Code)(httpError.Internal, httpError.Message)
			} else {
				genericError = gerr.New("unknown error", http.StatusInternalServerError)(err, nil)
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

	// TODO Server/Service Information
	//echoEngine.GET("/", func(c echo.Context) error {
	//	return c.JSONPretty(
	//		http.StatusOK,
	//		ServiceInfo{
	//			Name:       w.resource.ServiceName(),
	//			Identifier: w.resource.Identifier(),
	//		},
	//		" ",
	//	)
	//})

	if err := w.restApi.Init(echoEngine); err != nil {
		w.resource.Logger().Error("error on init http rest-server")
		return err
	}

	if w.restApi.Health() != nil {
		echoEngine.GET("/health", oculi.H(w.restApi.Health()))
	}
	return nil
}

func (w *WebServer) serve(sig chan os.Signal) {
	if err := w.resource.Echo().Start(fmt.Sprintf(":%d", w.resource.ServerPort())); err != nil {
		w.resource.Logger().Errorf("http rest-server interrupted %s", err.Error())
		sig <- syscall.SIGINT
	} else {
		w.resource.Logger().Info("starting apps")
	}
}

func (w *WebServer) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.ServerGracefullyDuration())
	defer cancel()

	if err := w.resource.Echo().Shutdown(ctx); err != nil {
		w.resource.Logger().Errorf("failed to shutdown http rest-server %s", err)
	}

	w.resource.Logger().Info("closing resource")
	if err := w.resource.Close(); err != nil {
		w.resource.Logger().Errorf("failed to close resource %s", err)
	}
}
