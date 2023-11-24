package httperror

import "net/http"

type (
	HttpError interface {
		error
		ResponseCode() int
		ResponseStatus() string
		Equal(err error) bool
		Metadata() any
		Source() error
		Detail() string
	}

	Seed func(source error, metadata any, args ...interface{}) HttpError
)

var ValidatorErrorHttpStatus = http.StatusUnprocessableEntity
