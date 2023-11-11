package tp

import (
	"strconv"

	"github.com/oculius/oculi/v2/common/error-extension"
)

type (
	boolParser struct{}
)

func BoolParser() Parser {
	return boolParser{}
}

func (b boolParser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}
	result, err := strconv.ParseBool(val)
	if err != nil {
		return nil, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	return result, nil
}
