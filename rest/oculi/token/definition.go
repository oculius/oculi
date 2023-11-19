package token

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
	errext "github.com/oculius/oculi/v2/common/http-error"
	tk "github.com/oculius/oculi/v2/rest/oculi/token/token-kind"
	ts "github.com/oculius/oculi/v2/rest/oculi/token/token-source"
)

type (
	Token interface {
		rawvalue() any

		Key() string
		IsRequired() bool
		Source() ts.TokenSource
		Type() tk.Kind
		Apply(ctx echo.Context) errext.HttpError
		String() string
	}

	ExtractTypeLimiter interface {
		bool | string |
			int | int8 | int16 | int32 | int64 |
			uint | uint8 | uint16 | uint32 | uint64 |
			float32 | float64 | time.Time | *multipart.FileHeader | []byte
	}
)
