package oculi

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/oculius/oculi/v2/common/error-extension"
	"net/http"
)

var (
	Translator        ut.Translator = nil
	ErrDataBinding                  = errext.New("data binding failed", http.StatusBadRequest)
	ErrDataValidation               = errext.New("data validation failed", http.StatusInternalServerError)
)
