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

	ErrorSeedBuilder interface {
		Build(source error, metadata any, args ...interface{}) Error
	}

	ErrorSeed func(source error, metadata any, args ...interface{}) Error
)

var ValidatorErrorHttpStatus = http.StatusUnprocessableEntity

func newSeed(seed ErrorSeedBuilder) ErrorSeed {
	return func(source error, metadata any, args ...interface{}) Error {
		return seed.Build(source, metadata, args...)
	}
}
