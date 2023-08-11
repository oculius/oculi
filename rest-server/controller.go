package rest

import (
	"bytes"
	"fmt"
	"github.com/oculius/oculi/v2/rest-server/oculi"
)

type (
	defaultCore struct {
		HealthModule
		components []Component
	}

	defaultComponent struct {
		path    string
		modules []Module
	}
)

func (r *defaultComponent) Init(engine oculi.Engine) error {
	buf := bytes.Buffer{}
	_, _ = fmt.Fprintf(&buf, "/%s", r.path)
	groupApi := engine.Group(buf.String())
	for _, mod := range r.modules {
		if err := mod.Init(groupApi); err != nil {
			return err
		}
	}
	return nil
}

func NewComponent(path string, modules ...Module) Component {
	return &defaultComponent{path, modules}
}

func (m *defaultCore) Init(engine oculi.Engine) error {
	for _, cmp := range m.components {
		if err := cmp.Init(engine); err != nil {
			return err
		}
	}
	return nil
}

func NewCore(health HealthModule, components ...Component) Core {
	return &defaultCore{health, components}
}
