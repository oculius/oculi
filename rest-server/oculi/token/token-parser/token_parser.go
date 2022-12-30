package tp

import (
	gerr "github.com/oculius/oculi/v2/common/error"
)

type (
	Parser interface {
		Parse(t Token, value any) (any, gerr.Error)
	}
)
