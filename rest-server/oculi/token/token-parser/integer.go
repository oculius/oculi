package tp

import (
	gerr "github.com/oculius/oculi/v2/common/error"
	"strconv"
)

type (
	intParser struct{}

	int8Parser struct{}

	int16Parser struct{}

	int32Parser struct{}

	int64Parser struct{}
)

func IntParser() Parser {
	return intParser{}
}

func (i intParser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	conVal, err := strconv.Atoi(val)
	if err != nil {
		return nil, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	return conVal, nil
}

func Int8Parser() Parser {
	return int8Parser{}
}

func (i int8Parser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericIntParser[int8](val, 8, t)
}

func Int16Parser() Parser {
	return int16Parser{}
}

func (i int16Parser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericIntParser[int16](val, 16, t)
}

func Int32Parser() Parser {
	return int32Parser{}
}

func (i int32Parser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericIntParser[int32](val, 32, t)
}

func Int64Parser() Parser {
	return int64Parser{}
}

func (i int64Parser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericIntParser[int64](val, 64, t)
}

func genericIntParser[T int8 | int16 | int32 | int64](val string, bitsize int, t Token) (T, gerr.Error) {
	var result T
	conVal, err := strconv.ParseInt(val, 10, bitsize)
	if err != nil {
		return result, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	result = T(conVal)
	return result, nil
}
