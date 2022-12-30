package tf

import (
	"github.com/labstack/echo/v4"
	gerr "github.com/oculius/oculi/v2/common/error"
	"mime/multipart"
)

type (
	Fetcher[T string | *multipart.FileHeader] interface {
		Fetch(echo.Context, Token) (T, gerr.Error)
	}

	Token interface {
		Key() string
		Metadata() any
		DataTypeString() string
	}
)
