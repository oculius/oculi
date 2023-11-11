package tp

import (
	"math/bits"
	"strconv"

	"github.com/oculius/oculi/v2/common/error-extension"
)

type (
	uintParser struct{}

	uint8Parser struct{}

	uint16Parser struct{}

	uint32Parser struct{}

	uint64Parser struct{}
)

func UintParser() Parser {
	return uintParser{}
}

func (i uintParser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	conVal, err := strconv.ParseUint(val, 10, bits.UintSize)
	if err != nil {
		return nil, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	return conVal, nil
}

func Uint8Parser() Parser {
	return uint8Parser{}
}

func (i uint8Parser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericUintParser[uint8](val, 8, t)
}

func Uint16Parser() Parser {
	return uint16Parser{}
}

func (i uint16Parser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericUintParser[uint16](val, 16, t)
}

func Uint32Parser() Parser {
	return uint32Parser{}
}

func (i uint32Parser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericUintParser[uint32](val, 32, t)
}

func Uint64Parser() Parser {
	return uint64Parser{}
}

func (i uint64Parser) Parse(t Token, value any) (any, errext.HttpError) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	return genericUintParser[uint64](val, 64, t)
}

func genericUintParser[T uint8 | uint16 | uint32 | uint64](val string, bitsize int, t Token) (T, errext.HttpError) {
	var result T
	conVal, err := strconv.ParseUint(val, 10, bitsize)
	if err != nil {
		return result, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	result = T(conVal)
	return result, nil
}
