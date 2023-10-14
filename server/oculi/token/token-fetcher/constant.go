package tf

import (
	"github.com/oculius/oculi/v2/common/error-extension"
	"net/http"
)

var (
	ErrFormFile        = errext.New("form file error", http.StatusInternalServerError)
	ErrCookie          = errext.New("unknown cookie error", http.StatusInternalServerError)
	ErrRequestNotFound = errext.New("http request is nil", http.StatusInternalServerError)
)
