package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/oculius/oculi/v2/utils/optional"
	"github.com/pkg/errors"
)

func LoadEnv() error {
	loadFileEnv := optional.Bool(os.Getenv("LOAD_CONFIG_FILE"), true)
	filename := os.Getenv("CONFIG_FILE_NAME")

	if filename == "" {
		filename = ".env"
	}

	if loadFileEnv {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("config: file '%s' not found", filename))
		}

		if err := godotenv.Load(filename); err != nil {
			return errors.Wrap(err, "failed to read from .env file")
		}
	}

	return nil
}

func NewEnv[T any](object *T) error {
	if err := LoadEnv(); err != nil {
		return err
	}

	if err := envconfig.Process("", object); err != nil {
		return errors.Wrap(err, "failed to read from env variable")
	}

	return nil
}
