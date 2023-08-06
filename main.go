package main

import (
	"fmt"
	di "github.com/oculius/oculi/v2/common/dependency-injection"
	bp_di "github.com/oculius/oculi/v2/common/dependency-injection/boilerplate"
	"github.com/oculius/oculi/v2/common/logs"
	ozap "github.com/oculius/oculi/v2/common/logs/zap"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/boilerplate"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"reflect"
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

var running = int64(0)
var count = 0

func (c C2) Init(route oculi.RouteGroup) error {
	apigroup := route.RGroup("/c2")
	apigroup.GET("/asd", func(ctx oculi.Context) error {
		id := count
		count++
		start := time.Now()
		go func() {
			for i := 0; i < 5; i++ {
				time.Sleep(time.Second)
				fmt.Println(id, ":", i)
			}
		}()
		resp := response.NewResponse("hai", map[string]any{"time": time.Now().Sub(start).String(), "start": start.Format(time.RFC3339Nano)}, true)
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
		resp := response.NewResponse("hola", nil, nil)
		return ctx.Send(resp)
	})
	return nil
}

var (
	_ rest.HealthController = HealthController{}
	_ rest.Component        = C1{}
	_ rest.Component        = C2{}
)

func (c Config) ServerGracefullyDuration() time.Duration {
	return time.Second * 5
}

func (c Config) IsDevelopmentMode() bool {
	return true
}

func (c HealthController) Health() oculi.HandlerFunc {
	return nil
}
func NewResource(l logs.Logger) rest.Resource {
	return bp_rest.Resource("123", 5555, l)
}

type A struct {
	N int
	M string
}

func printType(data any) {
	fmt.Println(reflect.ValueOf(data).Kind())
	fmt.Println(reflect.TypeOf(data).Kind())
}

func main() {
	//opts := memory.Options{
	//	Addr:     "localhost:6379",
	//	Password: "password",
	//}
	//
	//ctx := context.Background()
	//
	//m, _ := memory.NewRedis(json.New(), opts)
	//i := 0
	//for {
	//	select {
	//	case <-time.After(time.Millisecond * 250):
	//		fmt.Println(m.TTL(ctx, "x123"))
	//		i++
	//		if i > 20 {
	//			break
	//		}
	//	}
	//}
	//result := A{}
	//fmt.Println(m.Get(ctx, "gugu", &result))
	//fmt.Println(result)
	//result.N = 3
	//result.M = "agu"
	//fmt.Println(m.Set(ctx, "gugu", &result, 0))
	//fmt.Println(result)
	//var (
	//	i int
	//	f float32
	//)
	//fmt.Println(m.Get(ctx, "guguint", &i))
	//fmt.Println(m.Get(ctx, "gugufloat", &f))
	//fmt.Println(m.Get(ctx, "gugs", &f))
	//fmt.Println(i, f)

	var wg sync.WaitGroup
	time.Local, _ = time.LoadLocation("UTC")
	wg.Add(1)
	di.Register(
		bp_di.RestServer(),
		bp_di.APIVersion(1),
		fx.NopLogger,
		di.S(&wg),
		di.TS(Config{}, nil, nil, new(rest.Config)),
		di.TS(zap.AddCallerSkip(1), []string{`group:"log_opts"`}, new(zap.Option)),
		di.TP(ozap.NewProduction, []string{`group:"log_opts"`}, nil, nil),
		di.TS(C1{}, []string{`group:"v1_modules"`}, new(rest.Component)),
		di.TS(C2{}, []string{`group:"v1_modules"`}, new(rest.Component)),
		di.TS(HealthController{}, nil, new(rest.HealthController)),
		di.P(NewResource),
		fx.StartTimeout(time.Second*3),
	)
	deps := di.Dependencies()
	app := fx.New(deps...)

	app.Run()
	wg.Wait()
	//var wg sync.WaitGroup
	//wg.Add(3)
	//for i := 0; i < 3; i++ {
	//	go func(x int) {
	//		defer wg.Done()
	//		<-time.After(time.Second * 5)
	//		fmt.Printf("t%d done\n", x)
	//	}(i)
	//}
	wg.Wait()
}
