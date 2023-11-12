package authn

import (
	"net/http"

	errext "github.com/oculius/oculi/v2/common/http-error"
)

var (
	ErrAuthnFailedServer = errext.New("failed, server error", http.StatusInternalServerError)
	ErrAuthnFailedUser   = errext.New("failed, user error", http.StatusBadRequest)
)
