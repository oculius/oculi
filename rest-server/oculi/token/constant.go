package token

import (
	"fmt"
	gerr "github.com/oculius/oculi/v2/common/error"
	"github.com/pkg/errors"
	"net/http"
)

var (
	errRequiredValue = errors.New("missing required value")

	ErrInvalidSource   = gerr.New("invalid token source", http.StatusInternalServerError)
	ErrInvalidDataType = gerr.New("invalid token data type", http.StatusInternalServerError)
	ErrInvalidFetcher  = gerr.New("invalid token fetcher", http.StatusInternalServerError)(nil, nil)
	ErrRequiredValue   = gerr.NewConditional("token:required_value",
		func(i ...interface{}) string {
			return fmt.Sprintf("required %s:%s", i...)
		},
		func(error) int {
			return http.StatusBadRequest
		})
	ErrTypeCast = gerr.NewConditional("token:type_cast",
		func(i ...interface{}) string {
			return fmt.Sprintf("failed to cast '%v' to type %s", i...)
		},
		func(error) int {
			return http.StatusInternalServerError
		})
)
