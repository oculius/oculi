package tf

import (
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/error-extension"
	"github.com/pkg/errors"
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

func (f formfileFetcher) Fetch(ctx echo.Context, t Token) (*multipart.FileHeader, errext.HttpError) {
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

func (f formFetcher) Fetch(ctx echo.Context, t Token) (string, errext.HttpError) {
	return ctx.FormValue(t.Key()), nil
}
