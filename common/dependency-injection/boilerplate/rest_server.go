package bp_di

import (
	"context"
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

func newRestServer(mc rest.MainController, c rest.Config, res rest.Resource, lc fx.Lifecycle, wg *sync.WaitGroup) (rest.Server, error) {
	srv, err := rest.New(mc, res, c)
	if err != nil {
		return nil, err
	}
	lc.Append(newLCHook(wg, srv))
	return srv, nil
}

// Register Rest Server Provider & Invoker
// Required Dependencies: rest.HealthController, rest.Config, rest.Resource, *sync.WaitGroup, rest.Controller
func RestServer() di.ValuableComponent {
	opts := []fx.Option{
		di.P(newRestServer),
		di.I(func(srv rest.Server) {}),
		di.TP(bp_rest.MainController,
			[]string{
				`optional:"false"`,
				`group:"root_ctrl"`,
			},
			nil),
		di.TS("v1", []string{`name:"v1"`}),
		di.TP(bp_rest.RootController,
			[]string{
				`name:"v1"`,
				`group:"v1_ctrl"`,
			},
			[]string{
				`group:"root_ctrl"`,
			}),
	}
	return &holder{opts}
}
