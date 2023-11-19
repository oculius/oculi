package rest

import (
	"net/http"

	httperror "github.com/oculius/oculi/v2/common/http-error"
)

var (
	ErrNotFound         = httperror.New("not found", http.StatusNotFound)
	ErrMethodNotAllowed = httperror.New("not found", http.StatusMethodNotAllowed)
)
