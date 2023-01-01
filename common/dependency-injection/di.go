package di

import (
	"go.uber.org/fx"
	"reflect"
)

type (
	ValuableComponent interface {
		Dependencies() []fx.Option
	}

	Component interface {
		Child()
	}
)

func parse(item any, options *[]fx.Option) {
	callableComponent, ok := item.(Component)
	if ok {
		callableComponent.Child()
		return
	}

	component, ok := item.(ValuableComponent)
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

	*options = append(*options, P(item))
}

func Register(items ...any) {
	var opts []fx.Option

	i := Instance()

	for _, each := range items {
		parse(each, &opts)
	}

	if len(opts) > 0 {
		i.Add(opts)
	}
}

func Dependencies() []fx.Option {
	return Instance().Build()
}
