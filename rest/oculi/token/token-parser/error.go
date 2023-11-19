package tp

import (
	"fmt"
	"net/http"

	"github.com/oculius/oculi/v2/common/http-error"
)

var (
	ErrFormFile                = httperror.New("form file error", http.StatusInternalServerError)
	ErrInvalidInputValueString = httperror.New("invalid input type value for parsing",
		http.StatusInternalServerError,
	)(nil, map[string]any{"want": "string"})
	ErrInvalidInputValueFileHeader = httperror.New("invalid input type value for parsing",
		http.StatusInternalServerError,
	)(nil, map[string]any{"want": "file header"})
	ErrTypeParse = httperror.NewConditional("token_parser:type_parse",
		func(i ...interface{}) string {
			return fmt.Sprintf("failed to parse '%s' to type %s", i...)
		},
		func(error) int {
			return http.StatusBadRequest
		})
)
