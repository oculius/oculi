package rest

import (
	"bytes"
	"fmt"

	"github.com/oculius/oculi/v2/rest/oculi"
)

type (
	defaultComponent struct {
		path    string
		modules []Module
	}
)

func (r *defaultComponent) Init(engine *oculi.Engine) error {
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

func NewInternalComponent(path string, modules ...Module) InternalComponent {
	return &defaultComponent{path, modules}
}
