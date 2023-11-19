package rest

import (
	"time"

	"go.uber.org/fx"
)

type (
	Option struct {
		fx.In

		ServiceName              string        `name:"service_name"`
		ServicePort              int           `name:"service_port"`
		UpSince                  time.Time     `name:"up_since"`
		ShutdownGracefulDuration time.Duration `name:"shutdown_grace_period"`
		IsDevMode                bool          `name:"is_dev_mode"`
	}
)
