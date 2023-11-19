package tl

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	errext "github.com/oculius/oculi/v2/common/http-error"
	ts "github.com/oculius/oculi/v2/rest/oculi/token/token-source"
)

func loadString(source ts.TokenSource) (Loader[string], bool) {
	var router Loader[string]
	switch source {
	case ts.Query:
		router = Query()
	case ts.Parameter:
		router = URLParameter()
	case ts.Header:
		router = Header()
	case ts.Cookie:
		router = Cookie()
	case ts.Form:
		router = FormValue()
	default:
		return nil, false
	}
	return router, true
}

func loadFile(source ts.TokenSource) (Loader[*multipart.FileHeader], bool) {
	if source != ts.FormFile {
		return nil, false
	}
	return FormFile(), true
}

func Load(source ts.TokenSource, ctx echo.Context, token Token) (found bool, result any, err errext.HttpError) {
	found = false
	result = nil
	err = nil

	strVal, ok := loadString(source)
	if ok {
		found = true
		result, err = strVal.Load(ctx, token)
		return
	}
	fileVal, ok := loadFile(source)
	if ok {
		found = true
		result, err = fileVal.Load(ctx, token)
		return
	}
	return
}
