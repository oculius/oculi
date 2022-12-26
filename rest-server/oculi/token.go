package oculi

import (
	"fmt"
	"github.com/labstack/echo/v4"
	gerr "github.com/oculius/oculi/v2/common/error"
	tp "github.com/oculius/oculi/v2/rest-server/oculi/token-parser"
	"github.com/pkg/errors"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type (
	Kind        uint
	TokenSource uint

	token struct {
		key        string
		isRequired bool
		dataType   Kind
		source     TokenSource
		value      any
	}

	Token interface {
		Value() any
		Key() string
		IsRequired() bool
		Source() TokenSource
		Type() Kind
		Apply(ctx echo.Context) gerr.Error
	}
	TokenTypeLimiter interface {
		bool | string |
			int | int8 | int16 | int32 | int64 |
			uint | uint8 | uint16 | uint32 | uint64 |
			float32 | float64 | time.Time | *multipart.FileHeader | []byte
	}
)

var (
	errRequiredValue = errors.New("missing required value")
)

var (
	ErrFormFile        = gerr.New("form file error", http.StatusInternalServerError)
	ErrCookie          = gerr.New("unknown cookie error", http.StatusInternalServerError)
	ErrInvalidSource   = gerr.New("invalid token source", http.StatusInternalServerError)
	ErrInvalidDataType = gerr.New("invalid token data type", http.StatusInternalServerError)
	ErrRequestNotFound = gerr.New("http request is nil", http.StatusInternalServerError)
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

var tokenSourceString = []string{"invalid_source", "query", "parameter", "header", "cookie", "form", "form_file"}

const (
	InvalidSource TokenSource = iota
	Query
	Parameter
	Header
	Cookie
	Form
	FormFile
)

func (ts TokenSource) String() string {
	return tokenSourceString[ts]
}

var kindString = [...]string{
	"invalid_kind", "bool",
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64",
	"float32", "float64", "string", "time", "file_header",
	"uuid_string", "base36_string", "file_content_bytes", "file_content_base64",
}

const (
	InvalidKind Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Time
	FileHeader
	UUIDString
	Base36String
	FileContentBytes
	FileContentBase64
)

func (k Kind) String() string {
	return kindString[k]
}

func (k Kind) IsFromFormFile() bool {
	switch k {
	case FileHeader:
	case FileContentBytes:
	case FileContentBase64:
		return true
	}
	return false
}

// T stands for Token
func T(source TokenSource, key string, dataType Kind, isRequired bool) Token {
	return &token{
		key, isRequired, dataType, source, nil,
	}
}

func (t *token) Value() any {
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

func (t *token) getValue(ctx echo.Context) (value string, formFileValue *multipart.FileHeader, errResult gerr.Error) {
	errResult = nil
	formFileValue = nil
	value = ""
	switch t.source {
	case Query:
		value = ctx.QueryParam(t.key)
	case Parameter:
		value = ctx.Param(t.key)
	case Header:
		if ctx.Request() == nil {
			errResult = ErrRequestNotFound(nil, t.Metadata())
			return
		}
		value = ctx.Request().Header.Get(t.key)
	case Cookie:
		if ctx.Request() == nil {
			errResult = ErrRequestNotFound(nil, t.Metadata())
			return
		}
		cookie, err := ctx.Cookie(t.key)
		if err != nil && errors.Is(err, http.ErrNoCookie) {
			value = ""
			break
		} else if err != nil && !errors.Is(err, http.ErrNoCookie) {
			errResult = ErrCookie(err, t.Metadata())
			return
		}
		value = cookie.Value
	case Form:
		value = ctx.FormValue(t.key)
	case FormFile:
		if ctx.Request() == nil {
			errResult = ErrRequestNotFound(nil, t.Metadata())
			return
		}
		fh, err := ctx.FormFile(t.key)
		if err != nil && errors.Is(err, http.ErrMissingFile) {
			value = ""
			break
		} else if err != nil && !errors.Is(err, http.ErrMissingFile) {
			errResult = ErrFormFile(err, t.Metadata())
			return
		}
		formFileValue = fh
	default:
		errResult = ErrInvalidSource(nil, t.Metadata())
	}
	return
}

func (t *token) checkEmpty(val string, formFile *multipart.FileHeader) (bool, gerr.Error) {
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

func (t *token) parse(val string, ff *multipart.FileHeader) gerr.Error {
	var parser tp.Parser
	switch t.dataType {
	case Bool:
		parser = tp.BoolParser()
	case Int:
		parser = tp.IntParser()
	case Int8:
		parser = tp.Int8Parser()
	case Int16:
		parser = tp.Int16Parser()
	case Int32:
		parser = tp.Int32Parser()
	case Int64:
		parser = tp.Int64Parser()
	case Uint:
		parser = tp.UintParser()
	case Uint8:
		parser = tp.Uint8Parser()
	case Uint16:
		parser = tp.Uint16Parser()
	case Uint32:
		parser = tp.Uint32Parser()
	case Uint64:
		parser = tp.Uint64Parser()
	case Float32:
		parser = tp.Float32Parser()
	case Float64:
		parser = tp.Float64Parser()
	case String:
		parser = tp.StringParser()
	case Time:
		parser = tp.TimeParser()
	case FileHeader:
		parser = tp.FileHeaderParser()
	case UUIDString:
		parser = tp.UUIDStringParser()
	case Base36String:
		panic("parser is not implemented yet")
	case FileContentBytes:
		parser = tp.FileContentBytesParser()
	case FileContentBase64:
		parser = tp.FileContentBase64Parser()
	default:
		return ErrInvalidDataType(nil, t.Metadata())
	}

	var result any
	var err gerr.Error
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

func (t *token) Apply(ctx echo.Context) gerr.Error {
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

func ExtractTokenValue[T TokenTypeLimiter](token Token) (T, gerr.Error) {
	var result T
	val := token.Value()
	if val == nil {
		return result, nil
	}

	castedVal, ok := val.(T)
	if !ok {
		return result, ErrTypeCast(nil,
			map[string]any{
				"actual":   reflect.TypeOf(result).Kind().String(),
				"expected": token.Type().String(),
			}, val, token.Type().String())
	}
	return castedVal, nil
}
