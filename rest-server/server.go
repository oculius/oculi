package rest

import (
	"github.com/oculius/oculi/v2/common/error-extension"
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

	Initiable[T any] interface {
		Init(required T) error
	}

	Core interface {
		HealthController
		Module
	}

	Module    Initiable[oculi.Engine]
	Component Initiable[oculi.RouteGroup]

	//Component interface {
	//	Init(route oculi.RouteGroup) error
	//}

	Config interface {
		ServerGracefullyDuration() time.Duration
		IsDevelopmentMode() bool
	}

	HookFunction func(res Resource) error

	Option interface {
		Apply(w *webServer)
	}
)

var _ Server = &webServer{}

var (
	ErrNotFound         = errext.New("not found", http.StatusNotFound)
	ErrMethodNotAllowed = errext.New("not found", http.StatusMethodNotAllowed)
)
