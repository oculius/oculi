package tp

import (
	"fmt"
	"github.com/oculius/oculi/v2/common/error-extension"
	"net/http"
)

type (
	Integer interface {
		int | int8 | int16 | int32 | int64
	}

	UnsignedInteger interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	Number interface {
		Integer | UnsignedInteger
	}

	Parsable interface {
		Number
	}

	Token interface {
		Metadata() any
		DataTypeString() string
	}
)

var (
	ErrFormFile                = errext.New("form file error", http.StatusInternalServerError)
	ErrInvalidInputValueString = errext.New("invalid input type value for parsing",
		http.StatusInternalServerError,
	)(nil, map[string]any{"want": "string"})
	ErrInvalidInputValueFileHeader = errext.New("invalid input type value for parsing",
		http.StatusInternalServerError,
	)(nil, map[string]any{"want": "file header"})
	ErrTypeParse = errext.NewConditional("token_parser:type_parse",
		func(i ...interface{}) string {
			return fmt.Sprintf("failed to parse '%s' to type %s", i...)
		},
		func(error) int {
			return http.StatusBadRequest
		})
)
