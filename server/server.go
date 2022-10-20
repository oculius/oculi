package server

import (
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/logs"
	"time"
)

type (
	Engine interface {
		Run() error
		DevelopmentMode()

		BeforeRun(hf HookFunction) Engine
		AfterRun(hf HookFunction) Engine
		BeforeExit(hf HookFunction) Engine
		AfterExit(hf HookFunction) Engine
	}

	Resource interface {
		Echo() *echo.Echo
		ServiceName() string
		ServerPort() int
		Identifier() string
		Uptime() time.Time
		Logger() logs.Logger
		// Validator() validator.Validator
		Close() error
	}

	RestAPI interface {
		Register(ec *echo.Echo) error
		Health() echo.HandlerFunc
	}

	Config interface {
		ServerGracefullyDuration() time.Duration
		Instance() any
	}

	HookFunction func(res Resource) error

	WebServer struct {
		restApi        RestAPI
		resource       Resource
		config         Config
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

var _ Engine = &WebServer{}
