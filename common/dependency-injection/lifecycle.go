package di

import (
	"context"
	"github.com/oculius/oculi/v2/rest-server"
	"go.uber.org/fx"
	"golang.org/x/sys/unix"
	"sync"
)

func newLifecycleHook(wg *sync.WaitGroup, srv rest.Server) fx.Hook {
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
