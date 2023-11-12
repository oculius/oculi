package oculi

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/oculius/oculi/v2/common/http-error"
	"net/http"
)

var (
	Translator        ut.Translator = nil
	ErrDataBinding                  = httperror.New("data binding failed", http.StatusBadRequest)
	ErrDataValidation               = httperror.New("data validation failed", http.StatusInternalServerError)
)
