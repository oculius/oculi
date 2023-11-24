package rest

import (
	"bytes"
	"fmt"

	"github.com/oculius/oculi/v2/rest/oculi"
)

type (
	router struct {
		path    string
		modules []AccessPoint
	}
)

func (r *router) OnStart(engine *oculi.Engine) error {
	buf := bytes.Buffer{}
	_, _ = fmt.Fprintf(&buf, "/%s", r.path)
	groupApi := engine.Group(buf.String())
	for _, mod := range r.modules {
		if err := mod.OnStart(groupApi); err != nil {
			return err
		}
	}
	return nil
}

func NewGateway(path string, modules ...AccessPoint) Gateway {
	return &router{path, modules}
}
