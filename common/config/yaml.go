package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func NewYaml[T any](object *T, filename string) error {
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
