package di

import (
	"fmt"
	"go.uber.org/fx"
	"reflect"
	"sync"
	"time"
)

type (
	IndirectDependency interface {
		Dependencies() []fx.Option
	}

	Component interface {
		Child()
	}

	App struct {
		FxApp     *fx.App
		WaitGroup *sync.WaitGroup
	}
)

func parse(item any, options *[]fx.Option) {
	callableComponent, ok := item.(Component)
	if ok {
		callableComponent.Child()
		return
	}

	component, ok := item.(IndirectDependency)
	if ok {
		res := component.Dependencies()
		*options = append(*options, res...)
		return
	}

	opts, ok := item.([]fx.Option)
	if ok {
		*options = append(*options, opts...)
		return
	}

	opt, ok := item.(fx.Option)
	if ok {
		*options = append(*options, opt)
		return
	}

	if reflect.ValueOf(item).Kind() != reflect.Func {
		return
	}

	*options = append(*options, P(item))
}

func Compose(items ...any) {
	var opts []fx.Option

	i := newInstance()

	for _, each := range items {
		parse(each, &opts)
	}

	if len(opts) > 0 {
		i.Add(opts)
	}
}

var (
	isStartUpTimeoutSet = false
)

func Build() *App {
	var wg sync.WaitGroup

	buildTime := time.Now()

	if !isStartUpTimeoutSet {
		StartupTimeout(3 * time.Second)
	}

	deps := newInstance().Build()
	deps = append(deps, S(&wg))
	fmt.Printf("Took %s to build and prepare the application...\n", time.Now().Sub(buildTime).String())

	runTime := time.Now()
	defer func() {
		fmt.Printf("Took %s to run the application...\n", time.Now().Sub(runTime).String())
	}()
	return &App{fx.New(deps...), &wg}
}

func NoDependencyInjectionTracer() {
	i := newInstance()
	i.Add([]fx.Option{fx.NopLogger})
}

func StartupTimeout(v time.Duration) {
	i := newInstance()
	i.Add([]fx.Option{fx.StartTimeout(v)})
	isStartUpTimeoutSet = true
}

func StopTimeout(v time.Duration) {
	i := newInstance()
	i.Add([]fx.Option{fx.StopTimeout(v)})
}

func (a *App) Run() {
	a.FxApp.Run()
	a.WaitGroup.Wait()
}
