package bp

import (
	"context"
	"sync"

	"github.com/oculius/oculi/v2/server"
	"go.uber.org/fx"
	"golang.org/x/sys/unix"
)

func newLifecycleHook(wg *sync.WaitGroup, srv server.Server) fx.Hook {
	wg.Add(1)
	return fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				err := srv.Run()
				if err != nil {
					panic(err.Error())
				}
				wg.Done()
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			srv.Signal(unix.SIGTERM)
			return nil
		},
	}
}
