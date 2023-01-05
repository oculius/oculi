package bp_rest

import (
	"bytes"
	"fmt"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/oculi"
)

type (
	mainController struct {
		rest.HealthController
		controllers []rest.RootController
	}

	rootController struct {
		path        string
		controllers []rest.Controller
	}
)

func (r *rootController) Init(engine oculi.Engine) error {
	buf := bytes.Buffer{}
	_, _ = fmt.Fprintf(&buf, "/%s", r.path)
	groupApi := engine.Group(buf.String())
	for _, ctrl := range r.controllers {
		if err := ctrl.Init(groupApi); err != nil {
			return err
		}
	}
	return nil
}

func RootController(path string, controllers ...rest.Controller) rest.RootController {
	return &rootController{path, controllers}
}

func (m *mainController) Init(engine oculi.Engine) error {
	for _, ctrl := range m.controllers {
		if err := ctrl.Init(engine); err != nil {
			return err
		}
	}
	return nil
}

func MainController(health rest.HealthController, controllers ...rest.RootController) rest.MainController {
	return &mainController{health, controllers}
}
