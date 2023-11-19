package di

import (
	"reflect"

	"go.uber.org/fx"
)

func parse(item any, options *[]fx.Option) {
	component, ok := item.(Triggerable)
	if ok {
		component.Trigger()
		return
	}

	container, ok := item.(Container)
	if ok {
		res := container.Content()
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
		*options = append(*options, AsValue(item))
	} else {
		*options = append(*options, AsFunction(item))
	}
}
