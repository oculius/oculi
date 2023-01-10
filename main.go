package main

import (
	"fmt"
	di "github.com/oculius/oculi/v2/common/dependency-injection"
	"github.com/oculius/oculi/v2/common/dependency-injection/boilerplate"
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/common/logs/zap"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/boilerplate"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"sync"
	"time"
)

type (
	Resource struct {
		rest.Resource
		fx.In
	}

	Config struct{}

	HealthController struct{}

	C1 struct{}
	C2 struct{}
)

func (c C2) Init(route oculi.RouteGroup) error {
	apigroup := route.RGroup("/c2")
	apigroup.GET("/asd", func(ctx oculi.Context) error {
		resp := response.NewOkResponse("hai", Resource{}, true)
		return ctx.SendPretty(resp)
	})
	apigroup.RGroup("/baba").Bundle("", func(group oculi.RouteGroup) {
		group.POST("", func(ctx oculi.Context) error {
			return nil
		})
		group.TRACE("", func(ctx oculi.Context) error {
			return nil
		})
		group.DELETE("", func(ctx oculi.Context) error {
			return nil
		})
		group.PUT("", func(ctx oculi.Context) error {
			return nil
		})
		group.PATCH("/123", func(ctx oculi.Context) error {
			return nil
		})
		group.OPTIONS("", func(ctx oculi.Context) error {
			return nil
		})
		group.CONNECT("", func(ctx oculi.Context) error {
			return nil
		})
		group.HEAD("", func(ctx oculi.Context) error {
			return nil
		})
	})
	return nil
}

func (c C1) Init(route oculi.RouteGroup) error {
	apigroup := route.RGroup("/c1")
	apigroup.GET("/asd", func(ctx oculi.Context) error {
		fmt.Println("run")
		resp := response.NewOkResponse("hola", nil, nil)
		return ctx.Send(resp)
	})
	return nil
}

var (
	_ rest.HealthController = HealthController{}
	_ rest.Controller       = C1{}
	_ rest.Controller       = C2{}
)

func (c Config) ServerGracefullyDuration() time.Duration {
	return time.Second * 5
}

func (c HealthController) Health() oculi.HandlerFunc {
	return nil
}
func NewResource(l logs.Logger) rest.Resource {
	return bp_rest.Resource("123", 5555, l)
}

func main() {
	var wg sync.WaitGroup
	time.Local, _ = time.LoadLocation("UTC")
	wg.Add(1)
	di.Register(
		bp_di.RestServer(),
		fx.NopLogger,
		di.S(&wg),
		di.TS(Config{}, nil, nil, new(rest.Config)),
		di.TS(zap.AddCallerSkip(1), []string{`group:"log_opts"`}, new(zap.Option)),
		di.TP(ozap.NewProduction, []string{`group:"log_opts"`}, nil, nil),
		di.TS(C1{}, []string{`group:"v1_ctrl"`}, new(rest.Controller)),
		di.TS(C2{}, []string{`group:"v1_ctrl"`}, new(rest.Controller)),
		di.TS(HealthController{}, nil, new(rest.HealthController)),
		di.D(func(srv rest.Server) rest.Server {
			srv.AfterRun(func(res rest.Resource) error {
				res.Engine().Use(func(next oculi.HandlerFunc) oculi.HandlerFunc {
					return func(c oculi.Context) error {
						start := time.Now()
						err := next(c)
						fmt.Println(rest.NewPrinter().FormatRequest(c, start))
						return err
					}
				})
				return nil
			})
			return srv
		}),
		di.P(NewResource),
		fx.StartTimeout(time.Second*3),
	)
	deps := di.Dependencies()
	app := fx.New(deps...)

	app.Run()
	wg.Wait()
}
