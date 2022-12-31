package di

import "go.uber.org/fx"

type (
	function struct {
		fn        any
		resultTag string
		paramTag  string
	}
)

// P stands for Provider
func P(fn any) fx.Option {
	f := &function{fn, "", ""}
	return f.Build()
}

// I stands for Invoker
func I(fn any) fx.Option {
	f := &function{fn, "", ""}
	return f.Invoke()
}

// TP stands for Tagged Provider
func TP(fn any, resultTag string, paramTag string) fx.Option {
	f := &function{fn, resultTag, paramTag}
	return f.Build()
}

// D stands for Decorator
func D(fn any) fx.Option {
	f := &function{fn, "", ""}
	return f.Decorate()
}

func (f *function) Build() fx.Option {
	var annotations []fx.Annotation
	if len(f.resultTag) > 0 {
		annotations = append(annotations, fx.ResultTags(f.resultTag))
	}
	if len(f.paramTag) > 0 {
		annotations = append(annotations, fx.ParamTags(f.paramTag))
	}

	if len(annotations) == 0 {
		return fx.Provide(f.fn)
	}

	return fx.Provide(
		fx.Annotate(
			f.fn,
			annotations...,
		),
	)
}

func (f *function) Decorate() fx.Option {
	return fx.Decorate(
		f.fn,
	)
}

func (f *function) Invoke() fx.Option {
	return fx.Invoke(
		f.fn,
	)
}
