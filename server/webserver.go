package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/server/oculi"
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
		w.resource.Logger().Error("error on init http server")
		return err
	}

	echoEngine.GET("/health", w.restApi.Health())
	return nil
}

func (w *WebServer) serve(sig chan os.Signal) {
	if err := w.resource.Echo().Start(fmt.Sprintf(":%d", w.resource.ServerPort())); err != nil {
		w.resource.Logger().Errorf("http server interrupted %s", err.Error())
		sig <- syscall.SIGINT
	} else {
		w.resource.Logger().Info("starting apps")
	}
}

func (w *WebServer) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.ServerGracefullyDuration())
	defer cancel()

	if err := w.resource.Echo().Shutdown(ctx); err != nil {
		w.resource.Logger().Errorf("failed to shutdown http server %s", err)
	}

	w.resource.Logger().Info("closing resource")
	if err := w.resource.Close(); err != nil {
		w.resource.Logger().Errorf("failed to close resource %s", err)
	}
}
