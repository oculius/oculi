package di

import "go.uber.org/fx"

type (
	Container interface {
		Content() []fx.Option
	}

	Triggerable interface {
		Trigger()
	}

	Storage interface {
		Add(opts []fx.Option)
		Clear()
		Build() []fx.Option
	}
)
