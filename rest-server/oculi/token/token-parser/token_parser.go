package tp

import "github.com/oculius/oculi/v2/common/error-extension"

type (
	Parser interface {
		Parse(t Token, value any) (any, errext.Error)
	}
)
