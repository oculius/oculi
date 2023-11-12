package di

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/fx"
)

type (
	App struct {
		FxApp     *fx.App
		WaitGroup *sync.WaitGroup
	}
)

func Build() *App {
	var wg sync.WaitGroup

	buildTime := time.Now()

	if !isStartUpTimeoutSet {
		StartupTimeout(3 * time.Second)
	}

	deps := newInstance().Build()
	deps = append(deps, Supplier(&wg))
	fmt.Printf("Took %s to build and prepare the application...\n", time.Now().Sub(buildTime).String())

	runTime := time.Now()
	defer func() {
		fmt.Printf("Took %s to run the application...\n", time.Now().Sub(runTime).String())
	}()
	return &App{fx.New(deps...), &wg}
}

func (a *App) Run() {
	a.FxApp.Run()
	a.WaitGroup.Wait()
}
