package authtoken

import (
	errext "github.com/oculius/oculi/v2/common/error-extension"
	"net/http"
)

var (
	ErrInvalidToken = errext.New("invalid token", http.StatusUnauthorized)
	ErrFailedToSign = errext.New("failed to sign", http.StatusInternalServerError)
)
