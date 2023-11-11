package tp

import (
	"github.com/oculius/oculi/v2/common/error-extension"
	"strconv"
)

type (
	float32Parser struct{}

	float64Parser struct{}
)

func Float32Parser() Parser {
	return float32Parser{}
}

func (i float32Parser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericFloatParser[float32](val, 32, t)
}

func Float64Parser() Parser {
	return float64Parser{}
}

func (i float64Parser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericFloatParser[float64](val, 64, t)
}

func genericFloatParser[T float32 | float64](val string, bitsize int, t Token) (T, errext.HttpError) {
	var result T
	conVal, err := strconv.ParseFloat(val, bitsize)
	if err != nil {
		return result, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	result = T(conVal)
	return result, nil
}
