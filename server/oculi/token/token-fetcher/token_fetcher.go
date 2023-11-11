package tf

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/error-extension"
)

type (
	Fetcher[T string | *multipart.FileHeader] interface {
		Fetch(echo.Context, Token) (T, errext.HttpError)
	}

	Token interface {
		Key() string
		Metadata() any
		DataTypeString() string
	}
)
