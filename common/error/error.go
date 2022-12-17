package error

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
