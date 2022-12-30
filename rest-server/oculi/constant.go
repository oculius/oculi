package oculi

import (
	ut "github.com/go-playground/universal-translator"
	gerr "github.com/oculius/oculi/v2/common/error"
	"net/http"
)

var (
	Translator        ut.Translator = nil
	ErrDataBinding                  = gerr.New("data binding failed", http.StatusBadRequest)
	ErrDataValidation               = gerr.New("data validation failed", http.StatusInternalServerError)
)
