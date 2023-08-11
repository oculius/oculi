package di

import (
	"fmt"
	"github.com/oculius/oculi/v2/rest-server"
	"go.uber.org/fx"
	"reflect"
	"sync"
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

func newRestServer[X rest.Core, Y rest.Config, Z rest.IResource](core X, c Y, res Z, lc fx.Lifecycle, wg *sync.WaitGroup) (rest.Server, error) {
	srv, err := rest.New(core, res, c)
	if err != nil {
		return nil, err
	}
	lc.Append(newLifecycleHook(wg, srv))
	return srv, nil
}

// Register Rest Server Provider & Invoker
// Required Dependencies: rest.HealthModule, rest.Config, rest.IResource, *sync.WaitGroup, rest.Module
func RestServer[X rest.Core, Y rest.Config, Z rest.IResource]() IndirectDependency {
	opts := []fx.Option{
		P(newRestServer[X, Y, Z]),
		I(func(srv rest.Server) {}),
		TP(rest.NewCore,
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
			TP(rest.NewComponent,
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

func ComponentSupplier(moduleName string, module rest.Module) IndirectDependency {
	return &singleHolder{TS(module, Tag{fmt.Sprintf(`group:"%s_modules"`, moduleName)}, new(rest.Module))}
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

// ComponentProvider acts same as ComponentSupplier, but moduleProviderFn should return rest.Module
func ComponentProvider(moduleName string, moduleProviderFn any, paramTag Tag) IndirectDependency {
	moduleProviderFnValidator(moduleProviderFn)
	return &singleHolder{TP(moduleProviderFn, paramTag, Tag{fmt.Sprintf(`group:"%s_modules"`, moduleName)})}
}
