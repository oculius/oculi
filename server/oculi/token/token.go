package token

import (
	"fmt"
	"github.com/labstack/echo/v4"
	errext "github.com/oculius/oculi/v2/common/http-error"
	tf2 "github.com/oculius/oculi/v2/server/oculi/token/token-fetcher"
	tp2 "github.com/oculius/oculi/v2/server/oculi/token/token-parser"
	"mime/multipart"
	"reflect"
	"strings"
	"time"
)

type (
	token struct {
		key        string
		isRequired bool
		dataType   Kind
		source     TokenSource
		value      any
	}

	Token interface {
		rawvalue() any

		Key() string
		IsRequired() bool
		Source() TokenSource
		Type() Kind
		Apply(ctx echo.Context) errext.HttpError
		String() string
	}

	ExtractTypeLimiter interface {
		bool | string |
			int | int8 | int16 | int32 | int64 |
			uint | uint8 | uint16 | uint32 | uint64 |
			float32 | float64 | time.Time | *multipart.FileHeader | []byte
	}
)

// T stands for Token
func T(source TokenSource, key string, dataType Kind, isRequired bool) Token {
	return &token{
		key, isRequired, dataType, source, nil,
	}
}

func (t *token) String() string {
	return fmt.Sprintf("%s:%s=%+v", t.source.String(), t.key, t.value)
}

func (t *token) rawvalue() any {
	return t.value
}

func (t *token) Key() string {
	return t.key
}

func (t *token) IsRequired() bool {
	return t.isRequired
}

func (t *token) Source() TokenSource {
	return t.source
}

func (t *token) Type() Kind {
	return t.dataType
}

func (t *token) Metadata() any {
	return map[string]any{
		"key":         t.key,
		"source":      t.source.String(),
		"data_type":   t.dataType.String(),
		"is_required": t.isRequired,
	}
}

func (t *token) DataTypeString() string {
	return t.dataType.String()
}

func (t *token) getValue(ctx echo.Context) (string, *multipart.FileHeader, errext.HttpError) {
	var (
		fh         *multipart.FileHeader
		val        string
		strFetcher tf2.Fetcher[string]
		fhFetcher  tf2.Fetcher[*multipart.FileHeader]
	)
	switch t.source {
	case Query:
		strFetcher = tf2.QueryFetcher()
	case Parameter:
		strFetcher = tf2.URLParameterFetcher()
	case Header:
		strFetcher = tf2.HeaderFetcher()
	case Cookie:
		strFetcher = tf2.CookieFetcher()
	case Form:
		strFetcher = tf2.FormValueFetcher()
	case FormFile:
		fhFetcher = tf2.FormFileFetcher()
	default:
		return "", nil, ErrInvalidSource(nil, t.Metadata())
	}

	if strFetcher != nil {
		v, e := strFetcher.Fetch(ctx, t)
		if e != nil {
			return "", nil, e
		}
		fh = nil
		val = v
	} else if fhFetcher != nil {
		v, e := fhFetcher.Fetch(ctx, t)
		if e != nil {
			return "", nil, e
		}
		val = ""
		fh = v
	} else {
		return "", nil, ErrInvalidFetcher
	}
	return val, fh, nil
}

func (t *token) checkEmpty(val string, formFile *multipart.FileHeader) (bool, errext.HttpError) {
	if (t.dataType.IsFromFormFile() && formFile == nil) || (!t.dataType.IsFromFormFile() && len(val) == 0) {
		t.value = nil
		if t.isRequired {
			return true, ErrRequiredValue(errRequiredValue, t.Metadata(), t.source.String(), t.key)
		} else {
			return true, nil
		}
	}
	return false, nil
}

func (t *token) parse(val string, ff *multipart.FileHeader) errext.HttpError {
	var parser tp2.Parser
	switch t.dataType {
	case Bool:
		parser = tp2.BoolParser()
	case Int:
		parser = tp2.IntParser()
	case Int8:
		parser = tp2.Int8Parser()
	case Int16:
		parser = tp2.Int16Parser()
	case Int32:
		parser = tp2.Int32Parser()
	case Int64:
		parser = tp2.Int64Parser()
	case Uint:
		parser = tp2.UintParser()
	case Uint8:
		parser = tp2.Uint8Parser()
	case Uint16:
		parser = tp2.Uint16Parser()
	case Uint32:
		parser = tp2.Uint32Parser()
	case Uint64:
		parser = tp2.Uint64Parser()
	case Float32:
		parser = tp2.Float32Parser()
	case Float64:
		parser = tp2.Float64Parser()
	case String:
		parser = tp2.StringParser()
	case Time:
		parser = tp2.TimeParser()
	case FileHeader:
		parser = tp2.FileHeaderParser()
	case UUIDString:
		parser = tp2.UUIDStringParser()
	case Base36String:
		panic("parser is not implemented yet")
	case FileContentBytes:
		parser = tp2.FileContentBytesParser()
	case FileContentBase64:
		parser = tp2.FileContentBase64Parser()
	default:
		return ErrInvalidDataType(nil, t.Metadata())
	}

	var result any
	var err errext.HttpError
	if t.dataType.IsFromFormFile() {
		result, err = parser.Parse(t, ff)
	} else {
		result, err = parser.Parse(t, val)
	}
	if err != nil {
		return err
	}
	t.value = result
	return nil
}

func (t *token) Apply(ctx echo.Context) errext.HttpError {
	val, formFile, err := t.getValue(ctx)
	if err != nil {
		return err
	}
	val = strings.TrimSpace(val)

	empty, err := t.checkEmpty(val, formFile)
	if empty {
		return err // if empty, return err even-though it is nil, empty value can't be parsed
	}

	err = t.parse(val, formFile)
	if err != nil {
		return err
	}

	return nil
}

func TokenValue[T ExtractTypeLimiter](token Token) (T, errext.HttpError) {
	var result T
	val := token.rawvalue()
	if val == nil {
		return result, nil
	}

	castedVal, ok := val.(T)
	if !ok {
		return result, ErrTypeCast(nil,
			map[string]any{
				"actual":   reflect.TypeOf(val).Kind().String(),
				"expected": token.Type().String(),
			}, val, token.Type().String())
	}
	return castedVal, nil
}
