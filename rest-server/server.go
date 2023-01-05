package rest

import (
	gerr "github.com/oculius/oculi/v2/common/error"
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"net/http"
	"os"
	"time"
)

type (
	Server interface {
		Run() error
		Signal(signal os.Signal)

		BeforeRun(HookFunction) Server
		AfterRun(HookFunction) Server
		BeforeExit(HookFunction) Server
		AfterExit(HookFunction) Server
	}

	Resource interface {
		Engine() oculi.Engine
		ServiceName() string
		ServerPort() int
		Uptime() time.Time
		Logger() logs.Logger
		// Validator() validator.Validator
		Close() error
	}

	HealthController interface {
		Health() oculi.HandlerFunc
	}

	MainController interface {
		HealthController
		RootController
	}

	RootController interface {
		Init(engine oculi.Engine) error
	}

	Controller interface {
		Init(route oculi.RouteGroup) error
	}

	Config interface {
		ServerGracefullyDuration() time.Duration
	}

	HookFunction func(res Resource) error

	Option interface {
		Apply(w *webServer)
	}
)

var _ Server = &webServer{}

var (
	ErrNotFound         = gerr.New("not found", http.StatusNotFound)
	ErrMethodNotAllowed = gerr.New("not found", http.StatusMethodNotAllowed)
)
