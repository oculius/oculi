package gerr

import (
	"github.com/pkg/errors"
	"net/http"
)

type (
	staticErrorSeed struct {
		detail     string
		httpStatus int
	}

	staticError struct {
		source     error
		detail     string
		httpStatus int
		metadata   any
	}
)

func New(detail string, httpStatus int) ErrorSeed {
	return (&staticErrorSeed{detail, httpStatus}).Build
}

func (s *staticErrorSeed) Build(source error, metadata any, _ ...interface{}) Error {
	if source == nil {
		source = errors.New(s.detail)
	}
	return &staticError{
		source: source, metadata: metadata, detail: s.detail, httpStatus: s.httpStatus,
	}
}

func (s *staticError) Error() string {
	return s.source.Error()
}

func (s *staticError) Detail() string {
	return s.detail
}

func (s *staticError) ResponseCode() int {
	return s.httpStatus
}

func (s *staticError) ResponseStatus() string {
	return http.StatusText(s.httpStatus)
}

func (s *staticError) Equal(err Error) bool {
	_, ok := err.(*staticError)
	if !ok {
		return false
	}
	return s.httpStatus == err.ResponseCode() && s.source.Error() == err.Error() &&
		s.detail == err.Detail()
}

func (s *staticError) Metadata() any {
	return s.metadata
}

func (s *staticError) Source() error {
	return s.source
}
