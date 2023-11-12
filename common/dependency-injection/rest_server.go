package di

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/oculius/oculi/v2/server"
	"go.uber.org/fx"
)

type (
	holder struct {
		opts []fx.Option
	}

	singleHolder struct {
		opt fx.Option
	}
)

func (h *holder) Dependencies() []fx.Option {
	return h.opts
}

func (sh *singleHolder) Dependencies() []fx.Option {
	return []fx.Option{sh.opt}
}

func newRestServer[Core server.Core, Config server.Config, Resource server.IResource](
	core Core, config Config, res Resource,
	lc fx.Lifecycle, wg *sync.WaitGroup) (server.Server, error) {
	srv, err := server.New(core, res, config)
	if err != nil {
		return nil, err
	}
	lc.Append(newLifecycleHook(wg, srv))
	return srv, nil
}

// Register Rest Server Provider & Invoker
// Required Dependencies: rest.HealthModule, rest.Config, rest.IResource, *sync.WaitGroup, rest.Module
func RestServer[X server.Core, Y server.Config, Z server.IResource]() IndirectDependency {
	opts := []fx.Option{
		P(newRestServer[X, Y, Z]),
		I(func(srv server.Server) {}),
		TP(server.NewCore,
			[]string{
				`optional:"false"`,
				`group:"components"`,
			},
			nil),
	}
	return &holder{opts}
}

func NewComponent(name string) IndirectDependency {
	return &holder{
		opts: []fx.Option{
			// Supply Path Name
			TS(fmt.Sprintf("%s", name), Tag{fmt.Sprintf(`name:"%s"`, name)}),

			// Create API Version
			TP(server.NewComponent,
				[]string{
					fmt.Sprintf(`name:"%s"`, name),
					fmt.Sprintf(`group:"%s_modules"`, name),
				},
				[]string{
					`group:"components"`,
				}),
		},
	}
}

func APIVersion(version int) IndirectDependency {
	return NewComponent(fmt.Sprintf("v%d", version))
}

func SupplyModule(moduleName string, module server.Module) IndirectDependency {
	return &singleHolder{TS(module, Tag{fmt.Sprintf(`group:"%s_modules"`, moduleName)}, new(server.Module))}
}

var restModule = reflect.TypeOf((*server.Module)(nil)).Elem()

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
func ProvideModule(moduleName string, moduleProviderFn any, paramTag Tag) IndirectDependency {
	moduleProviderFnValidator(moduleProviderFn)
	return &singleHolder{TP(moduleProviderFn, paramTag, Tag{fmt.Sprintf(`group:"%s_modules"`, moduleName)})}
}
