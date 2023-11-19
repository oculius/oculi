package tp

import "github.com/oculius/oculi/v2/common/http-error"

type (
	Parser interface {
		Parse(t Token, value any) (any, httperror.HttpError)
	}

	Integer interface {
		int | int8 | int16 | int32 | int64
	}

	UnsignedInteger interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	Number interface {
		Integer | UnsignedInteger
	}

	Parsable interface {
		Number
	}

	Token interface {
		Metadata() any
		DataTypeString() string
	}
)
