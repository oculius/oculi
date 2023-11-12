package di

import "go.uber.org/fx"

type (
	IndirectDependency interface {
		Dependencies() []fx.Option
	}

	Component interface {
		Child()
	}
)
