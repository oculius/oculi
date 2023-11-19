package bp

import "go.uber.org/fx"

type (
	genericContainer []fx.Option
)

func (h genericContainer) Content() []fx.Option {
	return h
}
