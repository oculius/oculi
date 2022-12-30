package rest

import (
	"github.com/labstack/echo/v4"
	gerr "github.com/oculius/oculi/v2/common/error"
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"net/http"
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
		Uptime() time.Time
		Logger() logs.Logger
		// Validator() validator.Validator
		Close() error
	}

	RestAPI interface {
		Init(echoEngine *echo.Echo) error
		Health() oculi.HandlerFunc
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

var (
	ErrNotFound         = gerr.New("not found", http.StatusNotFound)
	ErrMethodNotAllowed = gerr.New("not found", http.StatusMethodNotAllowed)
)
