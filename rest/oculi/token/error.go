package token

import (
	"fmt"
	"net/http"

	"github.com/oculius/oculi/v2/common/http-error"
	"github.com/pkg/errors"
)

var (
	errRequiredValue = errors.New("missing required value")

	ErrInvalidSource   = httperror.New("invalid token source", http.StatusInternalServerError)
	ErrInvalidDataType = httperror.New("invalid token data type", http.StatusInternalServerError)
	ErrInvalidFetcher  = httperror.New("invalid token fetcher", http.StatusInternalServerError)(nil, nil)
	ErrRequiredValue   = httperror.NewConditional("token:required_value",
		func(i ...interface{}) string {
			return fmt.Sprintf("required %s:%s", i...)
		},
		func(error) int {
			return http.StatusBadRequest
		})
	ErrTypeCast = httperror.NewConditional("token:type_cast",
		func(i ...interface{}) string {
			return fmt.Sprintf("failed to cast '%v' to type %s", i...)
		},
		func(error) int {
			return http.StatusInternalServerError
		})
)
