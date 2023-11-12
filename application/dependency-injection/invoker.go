package di

import "go.uber.org/fx"

func Invoker(fn any) fx.Option {
	f := &function{fn, nil, nil, nil}
	return f.validate().invoke()
}
