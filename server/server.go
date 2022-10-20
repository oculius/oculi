package server

import (
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/logs"
	"time"
)

type (
	IServer interface {
		Run() error
		DevelopmentMode()

		BeforeRun(hf HookFunction) IServer
		AfterRun(hf HookFunction) IServer
		BeforeExit(hf HookFunction) IServer
		AfterExit(hf HookFunction) IServer
	}

	IResource interface {
		Echo() *echo.Echo
		ServiceName() string
		ServerPort() int
		Identifier() string
		Uptime() time.Time
		Logger() logs.ILogger
		// Validator() validator.Validator
		Close() error
	}

	IRestAPI interface {
		Register(ec *echo.Echo) error
		Health() echo.HandlerFunc
	}

	IConfig interface {
		ServerGracefullyDuration() time.Duration
		Instance() any
	}

	HookFunction func(res IResource) error

	WebServer struct {
		restApi        IRestAPI
		resource       IResource
		config         IConfig
		useDefaultGZip bool

		afterRun   []HookFunction
		beforeRun  []HookFunction
		beforeExit []HookFunction
		afterExit  []HookFunction
	}

	Option interface {
		Apply(w *WebServer)
	}
)

var _ IServer = &WebServer{}
