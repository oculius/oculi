package tf

import (
	"github.com/labstack/echo/v4"
	"github.com/oculius/oculi/v2/common/error-extension"
	"mime/multipart"
)

type (
	Fetcher[T string | *multipart.FileHeader] interface {
		Fetch(echo.Context, Token) (T, errext.Error)
	}

	Token interface {
		Key() string
		Metadata() any
		DataTypeString() string
	}
)
