package di

import "go.uber.org/fx"

func AsValue(item any) fx.Option {
	f := &function{item, nil, nil, nil}
	return f.validate().supply()
}

func AsTaggedValue(item any, resultTag Tag, asInterface ...interface{}) fx.Option {
	f := &function{item, resultTag, nil, asInterface}
	return f.validate().supply()
}
