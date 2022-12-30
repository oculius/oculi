package tf

import (
	gerr "github.com/oculius/oculi/v2/common/error"
	"net/http"
)

var (
	ErrFormFile        = gerr.New("form file error", http.StatusInternalServerError)
	ErrCookie          = gerr.New("unknown cookie error", http.StatusInternalServerError)
	ErrRequestNotFound = gerr.New("http request is nil", http.StatusInternalServerError)
)
