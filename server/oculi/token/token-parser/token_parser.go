package tp

import "github.com/oculius/oculi/v2/common/http-error"

type (
	Parser interface {
		Parse(t Token, value any) (any, httperror.HttpError)
	}
)
