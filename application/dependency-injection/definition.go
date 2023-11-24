package di

import "go.uber.org/fx"

type (
	Container interface {
		Content() []fx.Option
	}

	Triggerable interface {
		Trigger()
	}
)
