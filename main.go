package main

import (
	"encoding/json"
	"errors"
	"fmt"
	authtoken "github.com/oculius/oculi/v2/common/auth-token"
	"github.com/oculius/oculi/v2/common/authz"
	"github.com/oculius/oculi/v2/common/dependency-injection"
	"github.com/oculius/oculi/v2/common/logs"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest-server"
	"github.com/oculius/oculi/v2/rest-server/oculi"
	"go.uber.org/zap"
	"reflect"
	"time"
)

type (
	Resource struct {
		rest.IResource
	}

	Config struct {
		ServerName string
		ServerPort int
	}

	HealthController struct{}

	C1 struct{}
	C2 struct {
		res Resource
	}
)

var running = int64(0)
var count = 0

func NewC2(X string, r Resource, c Config) rest.Module {
	return C2{r}
}

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
		group.GET("", func(ctx oculi.Context) error {
			return errors.New("123")
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
	_ rest.HealthModule = HealthController{}
	_ rest.Module       = C1{}
	_ rest.Module       = C2{}
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
func NewResource(l logs.Logger, c Config) Resource {
	res := rest.Resource(c.ServerName, c.ServerPort, l)
	result := Resource{res}
	return result
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

	x, err := json.Marshal(authz.Permissions{
		authz.Permission{Resource: "123", Action: "write"},
		authz.Permission{Resource: "456", Action: "read"}})
	fmt.Println(string(x), err)
	var prm authz.Permissions
	fmt.Println(json.Unmarshal([]byte(`["123 writesss ","456111 read"]`), &prm))
	fmt.Println(prm)

	time.Local, _ = time.LoadLocation("Asia/Jakarta")
	di.Compose(
		di.RestServer[rest.Core, Config, Resource](),
		di.APIVersion(1),
		di.APIVersion(2),
		di.S(Config{"ASD", 5612}),
		di.TS(zap.AddCallerSkip(1), di.Tag{`group:"log_opts"`}, new(zap.Option)),
		di.TP(logs.NewZapProduction, di.Tag{`group:"log_opts"`}, nil, nil),
		di.SupplyModule("v1", C1{}),
		di.TS("asd", di.Tag{`s_name:"taga"`}),
		di.TS("def", di.Tag{`name:"tag2"`}),
		di.ProvideModule("v2", NewC2, di.Tag{`s_name:"taga"`}),
		//di.ComponentProvider("v1", func() (error, int, rest.Module) { return nil, 0, nil }, nil),
		//di.TP(NewC2, nil, di.Tag{`group:"v1_modules"`}),
		di.TS(HealthController{}, nil, new(rest.HealthModule)),
		di.P(NewResource),
	)
	//di.NoDependencyInjectionTracer()
	//app := di.Build()
	//app.Run()

	jwtEngine := authtoken.NewJWT[int]("baba123", authtoken.HS512)
	fmt.Println(jwtEngine.Encode(&authtoken.Claims[int]{Data: 123}, time.Second*10))
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTE3NDU1NDksIm5iZiI6MTY5MTc0NTUzOSwiaWF0IjoxNjkxNzQ1NTM5LCJkYXRhIjoxMjMsImlkZW50aWZpZXIiOiIifQ.-TXu9NJfB-x8ukKzX7hN7brFlgALwWneW5yuQ1Nz1nLaxcIbnYRmRMkyNRZ9IB9pBuTS_ZuMPW8BsEnJKWspdw"
	fmt.Println(jwtEngine.Decode(token))
	fmt.Println(jwtEngine.Validate(token))
	//var wg sync.WaitGroup
	//wg.Add(3)
	//for i := 0; i < 3; i++ {
	//	go func(x int) {
	//		defer wg.Done()
	//		<-time.After(time.Second * 5)
	//		fmt.Printf("t%d done\n", x)
	//	}(i)
	//}
}
