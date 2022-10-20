package server

import (
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/logs"
	"time"
)

type (
	ServerEngine interface {
		Run() error
		DevelopmentMode()

		BeforeRun(hf HookFunction) ServerEngine
		AfterRun(hf HookFunction) ServerEngine
		BeforeExit(hf HookFunction) ServerEngine
		AfterExit(hf HookFunction) ServerEngine
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

	RestAPIEngine interface {
		Register(ec *echo.Echo) error
		Health() echo.HandlerFunc
	}

	Config interface {
		ServerGracefullyDuration() time.Duration
		Instance() any
	}

	HookFunction func(res Resource) error

	WebServer struct {
		restApi        RestAPIEngine
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

var _ ServerEngine = &WebServer{}
