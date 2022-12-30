package tp

import (
	"github.com/gofrs/uuid"
	gerr "github.com/oculius/oculi/v2/common/error"
	"time"
)

type (
	stringParser struct{}

	uuidStringParser struct{}

	timeParser struct{}
)

func (s stringParser) Parse(_ Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}
	return val, nil
}

func StringParser() Parser {
	return stringParser{}
}

func UUIDStringParser() Parser {
	return uuidStringParser{}
}

func TimeParser() Parser {
	return timeParser{}
}

func (tps timeParser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}

	conVal, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	return conVal, nil
}

func (u uuidStringParser) Parse(t Token, value any) (any, gerr.Error) {
	val, ok := value.(string)
	if !ok {
		return nil, ErrInvalidInputValueString
	}
	convVal, err := uuid.FromString(val)
	if err != nil {
		return nil, ErrTypeParse(err, t.Metadata(), val, t.DataTypeString())
	}
	return convVal.String(), nil
}
