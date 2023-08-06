package bp_di

import (
	"context"
	"fmt"
	di "github.com/oculius/oculi/v2/common/dependency-injection"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/boilerplate"
	"go.uber.org/fx"
	"golang.org/x/sys/unix"
	"sync"
)

type (
	lc struct {
		wg  *sync.WaitGroup
		srv rest.Server
	}

	holder struct {
		opts []fx.Option
	}
)

func (h *holder) Dependencies() []fx.Option {
	return h.opts
}

func newLCHook(wg *sync.WaitGroup, srv rest.Server) fx.Hook {
	x := &lc{wg, srv}
	return fx.Hook{
		OnStart: x.OnStart,
		OnStop:  x.OnStop,
	}
}

func (l *lc) OnStart(_ context.Context) error {
	go func() {
		err := l.srv.Run()
		if err != nil {
			panic(err.Error())
		}
		l.wg.Done()
	}()
	return nil
}

func (l *lc) OnStop(_ context.Context) error {
	l.srv.Signal(unix.SIGTERM)
	return nil
}

func newRestServer(core rest.Core, c rest.Config, res rest.Resource, lc fx.Lifecycle, wg *sync.WaitGroup) (rest.Server, error) {
	srv, err := rest.New(core, res, c)
	if err != nil {
		return nil, err
	}
	lc.Append(newLCHook(wg, srv))
	return srv, nil
}

// Register Rest Server Provider & Invoker
// Required Dependencies: rest.HealthController, rest.Config, rest.Resource, *sync.WaitGroup, rest.Component
func RestServer() di.IndirectDependency {
	opts := []fx.Option{
		di.P(newRestServer),
		di.I(func(srv rest.Server) {}),
		di.TP(bp_rest.NewCore,
			[]string{
				`optional:"false"`,
				`group:"components"`,
			},
			nil),
	}
	return &holder{opts}
}

func NewComponent(name string) di.IndirectDependency {
	return &holder{
		opts: []fx.Option{
			// Supply Path Name
			di.TS(fmt.Sprintf("%s", name), []string{fmt.Sprintf(`path_name:"%s"`, name)}),

			// Create API Version
			di.TP(bp_rest.NewComponent,
				[]string{
					fmt.Sprintf(`path_name:"%s"`, name),
					fmt.Sprintf(`group:"%s_modules"`, name),
				},
				[]string{
					`group:"components"`,
				}),
		},
	}
}

func APIVersion(version int) di.IndirectDependency {
	return NewComponent(fmt.Sprintf("v%d", version))
}
