package pckg

import (
	di "github.com/oculius/oculi/v2/application/dependency-injection"
	"go.uber.org/fx"
)

type (
	PackageContainer []fx.Option
)

func (p PackageContainer) Content() []fx.Option {
	return p
}

var _ di.Container = PackageContainer{}
