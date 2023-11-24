package pckg_rest

import (
	"fmt"

	di "github.com/oculius/oculi/v2/application/dependency-injection"
	pckg "github.com/oculius/oculi/v2/application/dependency-injection/package"
	"github.com/oculius/oculi/v2/rest"
	"go.uber.org/fx"
)

func gatewayFactory(name string) fx.Option {
	return di.AsTaggedFunction(
		func(accessPoints ...rest.AccessPoint) rest.Gateway {
			return rest.NewGateway(name, accessPoints...)
		},
		[]string{
			fmt.Sprintf(GroupAccessPointFormat, name),
		},
		[]string{
			GroupGateway,
		},
	)
}

func Gateways(name ...string) di.Container {
	if len(name) == 0 || name == nil {
		panic("gateways factory called but no gateway registered")
	}
	result := pckg.PackageContainer{}
	for _, each := range name {
		result = append(result, gatewayFactory(each))
	}
	return result
}
