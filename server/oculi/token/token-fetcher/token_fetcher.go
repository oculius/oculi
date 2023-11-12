package tf

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/http-error"
)

type (
	Fetcher[T string | *multipart.FileHeader] interface {
		Fetch(echo.Context, Token) (T, httperror.HttpError)
	}

	Token interface {
		Key() string
		Metadata() any
		DataTypeString() string
	}
)
