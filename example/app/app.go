package app

import (
	"github.com/ravielze/oculi/app"
	"github.com/ravielze/oculi/example/di"
	"github.com/ravielze/oculi/example/infrastructures"
	"github.com/ravielze/oculi/example/resources"
	mw "github.com/ravielze/oculi/middleware/token"
	"github.com/ravielze/oculi/server"

	webserver "github.com/ravielze/oculi/server/echo"
	"go.uber.org/dig"
)

func Run() {
	invoker := func(container *dig.Container) error {
		return container.Invoke(func(i infrastructures.Component, r resources.Resource) error {
			s := webserver.New(i, r)
			if r.Config.IsDevelopment() {
				s.DevelopmentMode()
			}
			s.BeforeRun(func(res server.Resource) error {
				r := res.(resources.Resource)
				res.Echo().Use(mw.EchoMiddleware(r.Tokenizer.DecodeAccessHeader))
				return nil
			})
			return s.Run()
		})
	}

	if err := app.Run(di.Container, invoker); err != nil {
		panic(err)
	}
}
