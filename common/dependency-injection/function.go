package di

import (
	"github.com/oculius/optio/iterator"
	"go.uber.org/fx"
	"strings"
)

type (
	function struct {
		item        any
		resultTag   []string
		paramTag    []string
		asInterface []any
	}
)

// P stands for Provider
func P(fn any) fx.Option {
	f := &function{fn, nil, nil, nil}
	return f.validate().Build()
}

// I stands for Invoker
func I(fn any) fx.Option {
	f := &function{fn, nil, nil, nil}
	return f.validate().Invoke()
}

// TP stands for Tagged Provider
func TP(fn any, paramTag []string, resultTag []string, asInterface ...interface{}) fx.Option {
	f := &function{fn, resultTag, paramTag, asInterface}
	return f.validate().Build()
}

// D stands for Decorator
func D(fn any) fx.Option {
	f := &function{fn, nil, nil, nil}
	return f.validate().Decorate()
}

// S stands for Supplier
func S(item any) fx.Option {
	f := &function{item, nil, nil, nil}
	return f.validate().Supply()
}

// TS stands for Tagged Supplier
func TS(item any, resultTag []string, asInterface ...interface{}) fx.Option {
	f := &function{item, resultTag, nil, asInterface}
	return f.validate().Supply()
}

func (f *function) Build() fx.Option {
	ann := f.getAnnotations()
	if len(ann) == 0 {
		return fx.Provide(f.item)
	}

	return fx.Provide(
		fx.Annotate(
			f.item,
			ann...,
		),
	)
}

func (f *function) Decorate() fx.Option {
	return fx.Decorate(
		f.item,
	)
}

func (f *function) Invoke() fx.Option {
	return fx.Invoke(
		f.item,
	)
}

func (f *function) getAnnotations() []fx.Annotation {
	var annotations []fx.Annotation
	if len(f.resultTag) > 0 {
		annotations = append(annotations, fx.ResultTags(f.resultTag...))
	}
	if len(f.paramTag) > 0 {
		annotations = append(annotations, fx.ParamTags(f.paramTag...))
	}
	if f.asInterface != nil {
		annotations = append(annotations, fx.As(f.asInterface...))
	}
	return annotations
}

func (f *function) validate() *function {
	if f.item == nil {
		panic("illegal nil item/function detected")
	}
	if len(f.resultTag) > 0 {
		f.resultTag = iterator.Filter(f.resultTag, func(each string) bool {
			return len(strings.TrimSpace(each)) > 0
		})
	}
	if len(f.paramTag) > 0 {
		f.paramTag = iterator.Filter(f.paramTag, func(each string) bool {
			return len(strings.TrimSpace(each)) > 0
		})
	}
	if len(f.asInterface) > 0 {
		f.asInterface = iterator.Filter(f.asInterface, func(each any) bool {
			return each != nil
		})
	}
	return f
}

func (f *function) Supply() fx.Option {
	ann := f.getAnnotations()
	if len(ann) == 0 {
		return fx.Supply(f.item)
	}

	return fx.Supply(
		fx.Annotate(
			f.item, ann...,
		),
	)
}
