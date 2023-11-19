package di

import (
	"time"

	"go.uber.org/fx"
)

func Compose(items ...any) {
	var opts []fx.Option

	i := getInstance()

	for _, each := range items {
		parse(each, &opts)
	}

	if len(opts) > 0 {
		i.Add(opts)
	}
}

var (
	isStartUpTimeoutSet = false
)

func NoDependencyInjectionTracer() {
	i := getInstance()
	i.Add([]fx.Option{
		fx.NopLogger,
		fx.ErrorHook(diErrorLogger{}),
	})
}

func StartupTimeout(v time.Duration) {
	i := getInstance()
	i.Add([]fx.Option{fx.StartTimeout(v)})
	isStartUpTimeoutSet = true
}

func StopTimeout(v time.Duration) {
	i := getInstance()
	i.Add([]fx.Option{fx.StopTimeout(v)})
}
