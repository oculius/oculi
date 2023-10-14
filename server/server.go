package server

import (
	"github.com/oculius/oculi/v2/common/error-extension"
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/server/oculi"
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

	IResource interface {
		Engine() oculi.Engine
		ServiceName() string
		ServerPort() int
		Uptime() time.Time
		Logger() logs.Logger
		// Validator() validator.Validator
		Close() error
	}

	HealthModule interface {
		Health() oculi.HandlerFunc
	}

	Initiable[T any] interface {
		Init(required T) error
	}

	Core interface {
		HealthModule
		Initiable[oculi.Engine]
	}

	Component Initiable[oculi.Engine]
	Module    Initiable[oculi.RouteGroup]

	//Component interface {
	//	Init(route oculi.RouteGroup) error
	//}

	Config interface {
		ServerGracefullyDuration() time.Duration
		IsDevelopmentMode() bool
	}

	HookFunction func(res IResource) error

	Option interface {
		Apply(w *webServer)
	}
)

var _ Server = &webServer{}

var (
	ErrNotFound         = errext.New("not found", http.StatusNotFound)
	ErrMethodNotAllowed = errext.New("not found", http.StatusMethodNotAllowed)
)
