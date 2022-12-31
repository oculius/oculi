package resource

import (
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type (
	resource struct {
		serverName string
		serverPort int
		l          logs.Logger
		e          oculi.Engine
		t          time.Time
	}
)

func New(serverName string, serverPort int, l logs.Logger) rest.Resource {
	e := oculi.New()
	return &resource{
		serverName, serverPort, l, e, time.Now(),
	}
}

func (r *resource) Engine() oculi.Engine {
	return r.e
}

func (r *resource) Uptime() time.Time {
	return r.t
}

func (r *resource) Logger() logs.Logger {
	return r.l
}

func (r *resource) ServiceName() string {
	return r.serverName
}

func (r *resource) ServerPort() int {
	return r.serverPort
}

func (r *resource) Close() error {

	var errMessage = make([]string, 0)

	if err := r.Engine().Close(); err != nil {
		errMessage = append(errMessage, err.Error())
	}

	if len(errMessage) > 0 {
		return errors.New(strings.Join(errMessage, "\n"))
	}
	return nil
}
