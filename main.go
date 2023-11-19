package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/oculius/oculi/v2/application/dependency-injection"
	"github.com/oculius/oculi/v2/application/dependency-injection/boilerplate"
	logs2 "github.com/oculius/oculi/v2/application/logs"
	"github.com/oculius/oculi/v2/auth/authz"
	"github.com/oculius/oculi/v2/common/response"
	"github.com/oculius/oculi/v2/rest"
	"github.com/oculius/oculi/v2/rest/oculi"
	"go.uber.org/zap"
)

type (
	Config struct {
		DevMode bool `envconfig:"DEV_MODE"`
	}

	C1 struct{}
	C2 struct {
	}
)

var running = int64(0)
var count = 0

func NewC2(X string, c Config) rest.Module {
	return C2{}
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
	}, []oculi.MiddlewareFunc{}...)
	return nil
}

var (
	_ rest.Module = C1{}
	_ rest.Module = C2{}
)

type A struct {
	N int
	M string
}

func printType(data any) {
	fmt.Println(reflect.ValueOf(data).Kind())
	fmt.Println(reflect.TypeOf(data).Kind())
}
func main() {
	//c := Config{}
	//err := config.NewEnv(&c)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(c)
	//return
	// Schema
	//fields := graphql.Fields{
	//	"hello": &graphql.Field{
	//		Type: graphql.String,
	//		Args: map[string]*graphql.ArgumentConfig{
	//			"test": {
	//				Type:         graphql.String,
	//				DefaultValue: nil,
	//				Description:  "",
	//			},
	//		},
	//		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//			fmt.Printf("%v\n", p.Context)
	//			return p.Args["test"], nil
	//		},
	//	},
	//}
	//rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	//schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	//schema, err := graphql.NewSchema(schemaConfig)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//h := handler.New(&handler.Config{
	//	Schema:   &schema,
	//	Pretty:   true,
	//	GraphiQL: true,
	//})

	//fnValue := reflect.ValueOf(di.P)
	//fnPointer := runtime.FuncForPC(fnValue.Pointer())
	//fmt.Println(fnPointer.Name())
	//panic("!")
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
	//result.ToEchoMiddleware = "agu"
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

	di.Compose(
		//di.AsTaggedFunction(func() oculi.HandlerFunc {
		//	return func(ctx oculi.Context) error {
		//		return ctx.AutoSend(response.NewResponse("OK?", nil, nil))
		//	}
		//}, nil, di.Tag{`name:"healthcheck"`}),
		bp.RestServer[rest.Core](),
		bp.RestServerOption(),
		//di.AsValue(rest.NewOption("testing", 4512, true, 5*time.Second)),
		bp.APIVersion(1),
		bp.APIVersion(2),
		di.AsValue(Config{}),
		di.AsTaggedValue(zap.AddCallerSkip(1), di.Tag{`group:"log_opts"`}, new(zap.Option)),
		di.AsTaggedFunction(logs2.NewZapProduction, di.Tag{`group:"log_opts"`}, nil, nil),
		bp.SupplyModule("v1", C1{}),
		di.AsTaggedValue("asd", di.Tag{`s_name:"taga"`}),
		di.AsTaggedValue("def", di.Tag{`name:"tag2"`}),
		bp.ProvideModule("v2", NewC2, di.Tag{`s_name:"taga"`}),
		//di.ComponentProvider("v1", func() (error, int, rest.Module) { return nil, 0, nil }, nil),
		//di.TP(NewC2, nil, di.Tag{`group:"v1_modules"`}),
		//di.Invoker(func(resource Option) {
		//	fn := oculi.FromHttpHandler(h.ServeHTTP)
		//	fn2 := func(ctx oculi.Context) error {
		//		fmt.Printf("%+v\n", ctx.Request().Context().Value(ctx.Request().Context()))
		//		return fn(ctx)
		//	}
		//	resource.Engine().GET("/graphql", fn2)
		//	resource.Engine().POST("/graphql", fn2)
		//}),
	)
	di.NoDependencyInjectionTracer()
	app := di.Build()
	app.Run()

	//jwtEngine := authtoken.NewJWT[int]("baba123", authtoken.HS512)
	//fmt.Println(jwtEngine.Encode(&authtoken.Claims[int]{Data: 123}, time.Second*10))
	//token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTE3NDU1NDksIm5iZiI6MTY5MTc0NTUzOSwiaWF0IjoxNjkxNzQ1NTM5LCJkYXRhIjoxMjMsImlkZW50aWZpZXIiOiIifQ.-TXu9NJfB-x8ukKzX7hN7brFlgALwWneW5yuQ1Nz1nLaxcIbnYRmRMkyNRZ9IB9pBuTS_ZuMPW8BsEnJKWspdw"
	//fmt.Println(jwtEngine.Decode(token))
	//fmt.Println(jwtEngine.Validate(token))
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
