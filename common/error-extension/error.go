package errext

import "net/http"

type (
	Error interface {
		error
		ResponseCode() int
		ResponseStatus() string
		Equal(err error) bool
		Metadata() any
		Source() error
		Detail() string
	}

	ErrorSeed func(source error, metadata any, args ...interface{}) Error
)

var ValidatorErrorHttpStatus = http.StatusUnprocessableEntity
