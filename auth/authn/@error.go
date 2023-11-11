package authn

import (
	"net/http"

	errext "github.com/oculius/oculi/v2/common/error-extension"
)

var (
	ErrAuthnFailedServer = errext.New("failed, server error", http.StatusInternalServerError)
	ErrAuthnFailedUser   = errext.New("failed, user error", http.StatusBadRequest)
)
