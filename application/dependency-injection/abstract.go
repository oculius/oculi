package di

import (
	"strings"

	"github.com/oculius/optio/iterator"
	"go.uber.org/fx"
)

type (
	function struct {
		item        any
		resultTag   Tag
		paramTag    Tag
		asInterface []any
	}

	Tag []string
)

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

func (f *function) decorate() fx.Option {
	return fx.Decorate(
		f.item,
	)
}

func (f *function) invoke() fx.Option {
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

func (f *function) supply() fx.Option {
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
