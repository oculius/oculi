package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func parseBool(val string, def bool) bool {
	res, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return res
}

func New(object interface{}) error {
	useFileEnv := parseBool(os.Getenv("USE_CONFIG_FILE"), true)
	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", object); err != nil {
			return errors.Wrap(err, "failed to read from env variable")
		}
		return nil
	}
	if useFileEnv {
		if err := godotenv.Load(filename); err != nil {
			return errors.Wrap(err, "failed to read from .env file")
		}
	}

	if err := envconfig.Process("", object); err != nil {
		return errors.Wrap(err, "failed to read from env variable")
	}

	return nil
}
