package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func New(object interface{}) error {
	useFileEnv := !(strings.ToLower(os.Getenv("USE_CONFIG_FILE")) == "false")
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
