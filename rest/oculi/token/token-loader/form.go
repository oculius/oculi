package tl

import (
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/http-error"
	"github.com/pkg/errors"
)

type (
	formfile struct{}
	form     struct{}
)

func FormValue() Loader[string] {
	return form{}
}

func FormFile() Loader[*multipart.FileHeader] {
	return formfile{}
}

func (f formfile) Load(ctx echo.Context, t Token) (*multipart.FileHeader, httperror.HttpError) {
	if ctx.Request() == nil {
		return nil, ErrRequestNotFound(nil, t.Metadata())
	}
	fh, err := ctx.FormFile(t.Key())
	if err != nil && errors.Is(err, http.ErrMissingFile) {
		return nil, nil
	} else if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return nil, ErrFormFile(err, t.Metadata())
	}
	return fh, nil
}

func (f form) Load(ctx echo.Context, t Token) (string, httperror.HttpError) {
	return ctx.FormValue(t.Key()), nil
}
