package di

import (
	"go.uber.org/fx"
	"reflect"
)

type (
	Component interface {
		Construct()
	}
)

func Register(items ...any) {
	var opts []fx.Option

	i := Instance()

	for _, each := range items {
		wrapper, ok := each.(Component)
		if ok {
			wrapper.Construct()
			continue
		}

		fxopts, ok := each.(fx.Option)
		if ok {
			opts = append(opts, fxopts)
			continue
		}

		if reflect.ValueOf(each).Kind() != reflect.Func {
			continue
		}
		opts = append(opts, P(each))
	}

	if len(opts) > 0 {
		i.Add(opts)
	}
}
