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

	Initiable[T any] interface {
		Init(required T) error
	}

	Core interface {
		Initiable[*oculi.Engine]
		Healthcheck(ctx oculi.Context) error
	}

	ExternalComponent interface {
		Identifier() string
		Ping(ctx context.Context) error
	}

	InternalComponent Initiable[*oculi.Engine]
	Module            Initiable[oculi.RouteGroup]

	HookFunction func() error
)

var _ Server = &webServer{}
