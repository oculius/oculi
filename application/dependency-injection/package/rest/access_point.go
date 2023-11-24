package pckg_rest

import (
	"fmt"
	"reflect"

	di "github.com/oculius/oculi/v2/application/dependency-injection"
	pckg "github.com/oculius/oculi/v2/application/dependency-injection/package"
	"github.com/oculius/oculi/v2/rest"
)

func AccessPoint(targetGatewayName string, ap rest.AccessPoint) di.Container {
	return pckg.PackageContainer{
		di.AsTaggedValue(
			ap,
			di.Tag{fmt.Sprintf(GroupAccessPointFormat, targetGatewayName)},
			new(rest.AccessPoint),
		),
	}
}

var interfaceType = reflect.TypeOf((*rest.AccessPoint)(nil)).Elem()

func providerFunctionValidator(providerFn any) {
	fn := reflect.TypeOf(providerFn)
	if fn.Kind() != reflect.Func || fn.NumOut() <= 0 {
		panic("providerFn is not a provider function")
	}

	found := false
	for i := 0; i < fn.NumOut(); i++ {
		if fn.Out(i).Implements(interfaceType) {
			found = true
			break
		}
	}

	if !found {
		panic("providerFn is not a providing rest.AccessPoint")
	}
}

func AccessPointProvider(targetGatewayName string, providerFn any, paramTag di.Tag) di.Container {
	providerFunctionValidator(providerFn)
	return pckg.PackageContainer{
		di.AsTaggedFunction(
			providerFn,
			paramTag,
			di.Tag{
				fmt.Sprintf(GroupAccessPointFormat, targetGatewayName),
			},
		),
	}
}
