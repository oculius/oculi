package gerr

import "net/http"

type (
	Error interface {
		error
		ResponseCode() int
		ResponseStatus() string
		Equal(err Error) bool
		Metadata() any
		Source() error
		Detail() string
	}

	ErrorSeed interface {
		Build(source error, metadata any, args ...interface{}) Error
	}
)

var ValidatorErrorHttpStatus = http.StatusUnprocessableEntity
