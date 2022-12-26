package tp

import (
	gerr "github.com/oculius/oculi/v2/common/error"
	"github.com/pkg/errors"
	"strings"
)

type (
	boolParser struct{}
)

func BoolParser() Parser {
	return boolParser{}
}

var (
	errInvalidBool = errors.New("value is not true/false")
)

func (b boolParser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}
	if !(strings.EqualFold(val, "true") || strings.EqualFold(val, "false")) {
		return nil, ErrTypeParse(errInvalidBool, t.Metadata(), val, t.DataTypeString())
	}
	result := strings.EqualFold(val, "true")
	return result, nil
}
