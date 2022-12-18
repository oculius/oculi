package cmn_err

import "net/http"

type (
	GenericError interface {
		error
		ResponseCode() int
		ResponseStatus() string
		Equal(err GenericError) bool
		Metadata() any
		Source() error
		Detail() string
	}

	DynamicError interface {
		Build(source error, metadata any, args ...interface{}) GenericError
	}
)

var ValidatorErrorHttpStatus = http.StatusUnprocessableEntity
