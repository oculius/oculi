package di

import "go.uber.org/fx"

func Supplier(item any) fx.Option {
	f := &function{item, nil, nil, nil}
	return f.validate().supply()
}

func TaggedSupplier(item any, resultTag Tag, asInterface ...interface{}) fx.Option {
	f := &function{item, resultTag, nil, asInterface}
	return f.validate().supply()
}
