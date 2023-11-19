package di

import "go.uber.org/fx"

func AsFunction(fn any) fx.Option {
	f := &function{fn, nil, nil, nil}
	return f.validate().Build()
}

func AsTaggedFunction(fn any, paramTag Tag, resultTag Tag, asInterface ...interface{}) fx.Option {
	f := &function{fn, resultTag, paramTag, asInterface}
	return f.validate().Build()
}
