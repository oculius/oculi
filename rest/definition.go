package rest

import (
	"context"
	"os"

	"github.com/oculius/oculi/v2/rest/oculi"
)

type (
	Server interface {
		Run() error
		Signal(signal os.Signal)

		BeforeRun(...HookFunction)
		AfterRun(...HookFunction)
		BeforeExit(...HookFunction)
		AfterExit(...HookFunction)
	}

	Startable[T any] interface {
		OnStart(parent T) error
	}

	Core interface {
		Startable[*oculi.Engine]
		Healthcheck(ctx oculi.Context) error
	}

	ExternalComponent interface {
		Identifier() string
		Ping(ctx context.Context) error
	}

	Gateway     Startable[*oculi.Engine]
	AccessPoint Startable[oculi.RouteGroup]

	HookFunction func() error
)

var _ Server = &webServer{}
