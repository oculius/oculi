package tf

import (
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/error-extension"
	"github.com/pkg/errors"
	"mime/multipart"
	"net/http"
)

type (
	formfileFetcher struct{}
	formFetcher     struct{}
)

func FormValueFetcher() Fetcher[string] {
	return formFetcher{}
}

func FormFileFetcher() Fetcher[*multipart.FileHeader] {
	return formfileFetcher{}
}

func (f formfileFetcher) Fetch(ctx echo.Context, t Token) (*multipart.FileHeader, errext.Error) {
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

func (f formFetcher) Fetch(ctx echo.Context, t Token) (string, errext.Error) {
	return ctx.FormValue(t.Key()), nil
}
