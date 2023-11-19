package token

import (
	"reflect"

	errext "github.com/oculius/oculi/v2/common/http-error"
)

func Extract[T ExtractTypeLimiter](token Token) (T, errext.HttpError) {
	var result T
	val := token.rawvalue()
	if val == nil {
		return result, nil
	}

	castedVal, ok := val.(T)
	if !ok {
		return result, ErrTypeCast(nil,
			map[string]any{
				"actual":   reflect.TypeOf(val).Kind().String(),
				"expected": token.Type().String(),
			}, val, token.Type().String())
	}
	return castedVal, nil
}
