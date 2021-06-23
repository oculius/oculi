package validator

import (
	"reflect"
)

type (
	Validator interface {
		RegisterValidation(tag string, fn interface{})
		RegisterAlias(alias string, tags string)
		RegisterCustomTypeFunc(fn CustomTypeFunc, types ...interface{})
		RegisterStructValidation(fn interface{}, types ...interface{})
		InstallDefault()
		Validate(object interface{}) error
	}

	CustomTypeFunc func(field reflect.Value) interface{}
)
