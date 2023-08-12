package errext

import (
	"github.com/pkg/errors"
	"net/http"
)

type (
	dynamicErrorSeed struct {
		id         string
		detail     func(...interface{}) string
		httpStatus func(error) int
	}

	dynamicError struct {
		source     error
		detail     string
		httpStatus int
		metadata   any
		seed       dynamicErrorSeed
	}

	FormatFunc        func(...interface{}) string
	HttpStatusMapFunc func(error) int
)

func NewConditional(id string, formatter FormatFunc, conditionalHttpStatus HttpStatusMapFunc) ErrorSeed {
	return (&dynamicErrorSeed{
		id, formatter, conditionalHttpStatus,
	}).build
}

func (d dynamicErrorSeed) build(source error, metadata any, args ...interface{}) Error {
	if d.detail == nil || d.httpStatus == nil {
		panic("seed formatter or conditional http status is nil")
	}
	if source == nil {
		source = errors.New(d.detail(args...))
	}
	return &dynamicError{
		source,
		d.detail(args...),
		d.httpStatus(source),
		metadata,
		d,
	}
}

func (d *dynamicError) Error() string {
	return d.source.Error()
}

func (d *dynamicError) ResponseCode() int {
	return d.httpStatus
}

func (d *dynamicError) ResponseStatus() string {
	return http.StatusText(d.httpStatus)
}

func (d *dynamicError) Equal(err error) bool {
	casted, ok := err.(*dynamicError)
	if !ok {
		return false
	}
	return casted.seed.id == d.seed.id
}

func (d *dynamicError) Metadata() any {
	return d.metadata
}

func (d *dynamicError) Source() error {
	return d.source
}

func (d *dynamicError) Detail() string {
	return d.detail
}
