package gerr

import (
	"net/http"
	"reflect"
)

type (
	staticError struct {
		source     error
		detail     string
		httpStatus int
		metadata   any
	}
)

func NewError(source error, detail string, httpStatus int, metadata any) Error {
	return &staticError{source, detail, httpStatus, metadata}
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
		s.detail == err.Detail() && reflect.DeepEqual(s.metadata, err.Metadata())
}

func (s *staticError) Metadata() any {
	return s.metadata
}

func (s *staticError) Source() error {
	return s.source
}
