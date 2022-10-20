package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	w.resource.Echo().Use(middleware.Recover())
	// TODO Server/Validator
	// w.resource.Echo().Validator = w.resource.Validator()
	w.resource.Echo().Logger = w.resource.Logger()
	w.resource.Echo().Logger.SetLevel(log.INFO)

	if w.useDefaultGZip {
		w.resource.Echo().Use(middleware.Gzip())
	}

	//if w.useDefaultCors {
	//	w.resource.Echo().Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//		AllowOrigins: []string{"*"},
	//		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	//	}))
	//}

	// TODO Server/Custom Context
	//w.resource.Echo().Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		ctx, ok := c.(*oculiContext.Context)
	//		if !ok {
	//			ctx = oculiContext.New(c)
	//		}
	//		return next(ctx)
	//	}
	//})

	// TODO Server/Not Found Handler
	//echo.NotFoundHandler = func(c echo.Context) error {
	//	ctx := oculiContext.New(c)
	//	ctx.AddError(http.StatusNotFound, errors.New("not found"))
	//	return ctx.JSONPretty(
	//		ctx.ResponseCode(),
	//		NotFoundStruct{
	//			Code:    ctx.ResponseCode(),
	//			Message: ctx.Errors()[0].Error(),
	//		},
	//		" ",
	//	)
	//}

	// TODO Server/Method Not Allowed
	//echo.MethodNotAllowedHandler = func(c echo.Context) error {
	//	ctx := oculiContext.New(c)
	//	ctx.AddError(http.StatusMethodNotAllowed, errors.New("method not allowed"))
	//	return ctx.JSONPretty(
	//		ctx.ResponseCode(),
	//		NotFoundStruct{
	//			Code:    ctx.ResponseCode(),
	//			Message: ctx.Errors()[0].Error(),
	//		},
	//		" ",
	//	)
	//}

	// TODO Server/Service Information
	//w.resource.Echo().GET("/", func(c echo.Context) error {
	//	return c.JSONPretty(
	//		http.StatusOK,
	//		ServiceInfo{
	//			Name:       w.resource.ServiceName(),
	//			Identifier: w.resource.Identifier(),
	//		},
	//		" ",
	//	)
	//})

	if err := w.restApi.Register(w.resource.Echo()); err != nil {
		// TODO Server/Logger Error Log
		// w.resource.Logger().Error("error on register http")
		return err
	}

	w.resource.Identifier()
	w.resource.Echo().GET("/health", w.restApi.Health())
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
