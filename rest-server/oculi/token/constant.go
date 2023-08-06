package token

import (
	"fmt"
	"github.com/oculius/oculi/v2/common/error-extension"
	"github.com/pkg/errors"
	"net/http"
)

var (
	errRequiredValue = errors.New("missing required value")

	ErrInvalidSource   = errext.New("invalid token source", http.StatusInternalServerError)
	ErrInvalidDataType = errext.New("invalid token data type", http.StatusInternalServerError)
	ErrInvalidFetcher  = errext.New("invalid token fetcher", http.StatusInternalServerError)(nil, nil)
	ErrRequiredValue   = errext.NewConditional("token:required_value",
		func(i ...interface{}) string {
			return fmt.Sprintf("required %s:%s", i...)
		},
		func(error) int {
			return http.StatusBadRequest
		})
	ErrTypeCast = errext.NewConditional("token:type_cast",
		func(i ...interface{}) string {
			return fmt.Sprintf("failed to cast '%v' to type %s", i...)
		},
		func(error) int {
			return http.StatusInternalServerError
		})
)
