package rbac

import (
	"net/http"

	errext "github.com/oculius/oculi/v2/common/error-extension"
)

var (
	ErrAuthorizationService = errext.New("authorization service error", http.StatusInternalServerError)
	ErrInvalidResourceName  = errext.New("invalid resource name", http.StatusBadRequest)
	ErrInvalidActionName    = errext.New("invalid action name", http.StatusBadRequest)
	ErrForbidden            = errext.New("no permission", http.StatusForbidden)
)
