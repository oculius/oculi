package pckg_rest

import (
	"sync"

	di "github.com/oculius/oculi/v2/application/dependency-injection"
	pckg "github.com/oculius/oculi/v2/application/dependency-injection/package"

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
// Required Dependencies: rest.HealthModule, rest.Config, *sync.WaitGroup, rest.AccessPoint
func Server[Core rest.Core]() di.Container {
	return pckg.PackageContainer{
		di.AsFunction(oculi.New),
		di.AsFunction(newRestServer[Core]),
		di.Invoker(func(srv rest.Server) {}),
		di.AsTaggedFunction(rest.NewCore,
			[]string{
				`optional:"true" ` + NameHealthcheck,
				`optional:"false" ` + GroupGateway,
				`optional:"false" ` + GroupExternals,
			},
			nil),
	}
}
