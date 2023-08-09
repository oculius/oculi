package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
)

func parseBool(val string, def bool) bool {
	res, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return res
}

func NewYaml(object any, filename string) error {
	if fileInfo, err := os.Stat(filename); os.IsNotExist(err) {
		return err
	} else if !strings.HasSuffix(fileInfo.Name(), ".yml") &&
		!strings.HasSuffix(fileInfo.Name(), ".yaml") {
		return errors.New(fmt.Sprintf("config: file '%s' is not yml/yaml", filename))
	}
	buff, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read yaml file")
	}

	err = yaml.Unmarshal(buff, object)
	if err != nil {
		return errors.Wrap(err, "failed to parse file data to yaml")
	}
	return nil
}

func NewEnv(object interface{}) error {
	useFileEnv := parseBool(os.Getenv("USE_CONFIG_FILE"), true)
	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err2 := envconfig.Process("", object); err2 != nil {
			return errors.Wrap(err2, "failed to read from env variable")
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
