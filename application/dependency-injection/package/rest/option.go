package pckg_rest

import (
	"fmt"
	"os"
	"strconv"
	"time"

	di "github.com/oculius/oculi/v2/application/dependency-injection"
	pckg "github.com/oculius/oculi/v2/application/dependency-injection/package"
	"github.com/oculius/oculi/v2/utils/optional"
)

func getEnvOrDef(name string, def string) string {
	v, found := os.LookupEnv(name)
	if !found {
		return def
	}
	return v
}

func getPort() int {
	fmt.Println(getEnvOrDef("SERVICE_PORT", "8001"))
	port, err := strconv.Atoi(getEnvOrDef("SERVICE_PORT", "8001"))
	if err != nil {
		return 8001
	}
	return port
}

func getShutdownGracePeriod() time.Duration {
	dur, err := time.ParseDuration(getEnvOrDef("SHUTDOWN_GRACE_PERIOD", "5s"))
	if err != nil {
		return 5 * time.Second
	}
	return dur
}

func getName() string {
	return getEnvOrDef("SERVICE_NAME", "Unnamed Service")
}

func isDevMode() bool {
	return optional.Bool(getEnvOrDef("DEV_MODE", "false"), false)
}

func RestServerOption() di.Container {
	return pckg.PackageContainer{
		di.AsTaggedFunction(
			getName,
			nil,
			di.Tag{`name:"service_name"`},
		),
		di.AsTaggedFunction(
			getPort,
			nil,
			di.Tag{`name:"service_port"`},
		),
		di.AsTaggedFunction(
			time.Now,
			nil,
			di.Tag{`name:"up_since"`},
		),
		di.AsTaggedFunction(
			getShutdownGracePeriod,
			nil,
			di.Tag{`name:"shutdown_grace_period"`},
		),
		di.AsTaggedFunction(
			isDevMode,
			nil,
			di.Tag{`name:"is_dev_mode"`},
		),
	}
}
