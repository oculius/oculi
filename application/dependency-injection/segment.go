package di

import "go.uber.org/fx"

type (
	segment struct {
		child []fx.Option
		name  string
	}
)

func Segment(name string, fn ...any) fx.Option {
	var opts []fx.Option
	for _, each := range fn {
		parse(each, &opts)
	}
	m := segment{opts, name}
	return m.build()
}

func (m *segment) build() fx.Option {
	X := len(m.child)
	result := make([]fx.Option, X)
	i := 0
	for _, each := range m.child {
		if each == nil {
			continue
		}
		result[i] = each
		i++
	}
	return fx.Module(m.name, result...)
}
