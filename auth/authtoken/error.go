package authtoken

import (
	errext "github.com/oculius/oculi/v2/common/http-error"
	"net/http"
)

var (
	ErrInvalidToken = errext.New("invalid token", http.StatusUnauthorized)
	ErrFailedToSign = errext.New("failed to sign", http.StatusInternalServerError)
)
