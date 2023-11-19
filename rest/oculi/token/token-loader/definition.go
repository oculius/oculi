package tl

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/http-error"
)

type (
	Loader[T string | *multipart.FileHeader] interface {
		Load(echo.Context, Token) (T, httperror.HttpError)
	}

	Token interface {
		Key() string
		Metadata() any
		DataTypeString() string
	}
)
