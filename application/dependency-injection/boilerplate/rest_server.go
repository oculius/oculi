package bp

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/oculius/oculi/v2/application/dependency-injection"
	"github.com/oculius/oculi/v2/application/logs"
	"github.com/oculius/oculi/v2/rest"
	"github.com/oculius/oculi/v2/rest/oculi"
	"go.uber.org/fx"
)

func newRestServer[Core rest.Core](
	core Core,
	opt rest.Option,
	lc fx.Lifecycle,
	engine *oculi.Engine,
	logger logs.Logger,
	wg *sync.WaitGroup,
) (rest.Server, error) {
	l := logger.With("serviceName", opt.ServiceName)
	srv, err := rest.New(core, opt, engine, l)
	if err != nil {
		return nil, err
	}
	lc.Append(newLifecycleHook(wg, srv))
	return srv, nil
}

// Register Rest Server Provider & Invoker
// Required Dependencies: rest.HealthModule, rest.Config, *sync.WaitGroup, rest.Module
func RestServer[X rest.Core]() di.Container {
	return genericContainer{
		di.AsFunction(oculi.New),
		di.AsFunction(newRestServer[X]),
		di.Invoker(func(srv rest.Server) {}),
		di.AsTaggedFunction(rest.NewCore,
			[]string{
				`optional:"true" name:"healthcheck"`,
				`optional:"false" group:"internals"`,
				`optional:"false" group:"externals"`,
			},
			nil),
	}
}

func NewInternalComponents(name string) di.Container {
	return genericContainer{
		di.AsTaggedFunction(func(modules ...rest.Module) rest.InternalComponent {
			return rest.NewInternalComponent(name, modules...)
		},
			[]string{
				fmt.Sprintf(`group:"%s_modules"`, name),
			},
			[]string{
				`group:"internals"`,
			}),
	}
}

func APIVersion(version int) di.Container {
	return NewInternalComponents(fmt.Sprintf("v%d", version))
}

func SupplyModule(moduleName string, module rest.Module) di.Container {
	return genericContainer{di.AsTaggedValue(module, di.Tag{fmt.Sprintf(`group:"%s_modules"`, moduleName)}, new(rest.Module))}
}

var restModule = reflect.TypeOf((*rest.Module)(nil)).Elem()

func moduleProviderFnValidator(moduleProviderFn any) {
	fn := reflect.TypeOf(moduleProviderFn)
	if fn.Kind() != reflect.Func || fn.NumOut() <= 0 {
		panic("moduleProviderFn is not a provider function")
	}

	found := false
	for i := 0; i < fn.NumOut(); i++ {
		if fn.Out(i).Implements(restModule) {
			found = true
			break
		}
	}

	if !found {
		panic("moduleProviderFn is not a providing rest.Module")
	}
}

// ProvideModule acts same as SupplyModule, but moduleProviderFn should return rest.Module
func ProvideModule(moduleName string, moduleProviderFn any, paramTag di.Tag) di.Container {
	moduleProviderFnValidator(moduleProviderFn)
	return genericContainer{
		di.AsTaggedFunction(
			moduleProviderFn,
			paramTag,
			di.Tag{fmt.Sprintf(`group:"%s_modules"`, moduleName)},
		),
	}
}
