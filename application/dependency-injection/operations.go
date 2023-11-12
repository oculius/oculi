package di

import (
	"reflect"
	"time"

	"go.uber.org/fx"
)

func Compose(items ...any) {
	var opts []fx.Option

	i := newInstance()

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
	i := newInstance()
	i.Add([]fx.Option{
		fx.NopLogger,
		fx.ErrorHook(diErrorLogger{}),
	})
}

func StartupTimeout(v time.Duration) {
	i := newInstance()
	i.Add([]fx.Option{fx.StartTimeout(v)})
	isStartUpTimeoutSet = true
}

func StopTimeout(v time.Duration) {
	i := newInstance()
	i.Add([]fx.Option{fx.StopTimeout(v)})
}

func parse(item any, options *[]fx.Option) {
	callableComponent, ok := item.(Component)
	if ok {
		callableComponent.Child()
		return
	}

	component, ok := item.(IndirectDependency)
	if ok {
		res := component.Dependencies()
		*options = append(*options, res...)
		return
	}

	opts, ok := item.([]fx.Option)
	if ok {
		*options = append(*options, opts...)
		return
	}

	opt, ok := item.(fx.Option)
	if ok {
		*options = append(*options, opt)
		return
	}

	if reflect.ValueOf(item).Kind() != reflect.Func {
		return
	}

	*options = append(*options, Provider(item))
}
