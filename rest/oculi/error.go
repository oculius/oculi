package oculi

import (
	"net/http"

	httperror "github.com/oculius/oculi/v2/common/http-error"
)

var (
	ErrDataBinding    = httperror.New("data binding failed", http.StatusBadRequest)
	ErrDataValidation = httperror.New("data validation failed", http.StatusInternalServerError)
)
