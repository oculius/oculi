package token

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/labstack/echo/v4"
	errext "github.com/oculius/oculi/v2/common/http-error"
	tk "github.com/oculius/oculi/v2/rest/oculi/token/token-kind"
	tl "github.com/oculius/oculi/v2/rest/oculi/token/token-loader"
	tp "github.com/oculius/oculi/v2/rest/oculi/token/token-parser"
	ts "github.com/oculius/oculi/v2/rest/oculi/token/token-source"
)

type (
	token struct {
		key        string
		isRequired bool
		dataType   tk.Kind
		source     ts.TokenSource
		value      any
	}
)

func New(source ts.TokenSource, key string, dataType tk.Kind, isRequired bool) Token {
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

func (t *token) Source() ts.TokenSource {
	return t.source
}

func (t *token) Type() tk.Kind {
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
		fh  *multipart.FileHeader = nil
		val                       = ""
	)
	found, loadedValue, err := tl.Load(t.source, ctx, t)
	if !found {
		return "", nil, ErrInvalidSource(nil, t.Metadata())
	}
	if err != nil {
		return "", nil, err
	}

	if strVal, okStr := loadedValue.(string); okStr {
		val = strVal
	} else if fileVal, okFile := loadedValue.(*multipart.FileHeader); okFile {
		fh = fileVal
	} else if !okStr && !okFile {
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
	parser, ok := tp.Get(t.dataType)
	if !ok || parser == nil {
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
