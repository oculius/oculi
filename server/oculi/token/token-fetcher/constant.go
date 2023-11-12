package tf

import (
	"github.com/oculius/oculi/v2/common/http-error"
	"net/http"
)

var (
	ErrFormFile        = httperror.New("form file error", http.StatusInternalServerError)
	ErrCookie          = httperror.New("unknown cookie error", http.StatusInternalServerError)
	ErrRequestNotFound = httperror.New("http request is nil", http.StatusInternalServerError)
)
