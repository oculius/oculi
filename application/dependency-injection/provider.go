package di

import "go.uber.org/fx"

func Provider(fn any) fx.Option {
	f := &function{fn, nil, nil, nil}
	return f.validate().Build()
}

func TaggedProvider(fn any, paramTag Tag, resultTag Tag, asInterface ...interface{}) fx.Option {
	f := &function{fn, resultTag, paramTag, asInterface}
	return f.validate().Build()
}
