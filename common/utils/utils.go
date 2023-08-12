package utils

import (
	"strings"
	"time"
)

type (
	Integer interface {
		int | int8 | int16 | int32 | int64
	}

	UnsignedInteger interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	Number interface {
		Integer | UnsignedInteger
	}
)

func PointerFactory[T any](val T) *T {
	i := new(T)
	*i = val
	return i
}

func ZeroOrNil[T Number | time.Duration](val *T) *T {
	if val == nil || *val == T(0) {
		return nil
	}
	return val
}

func EmptyOrNil[T string | time.Time](val *T) *T {
	if casted, ok := any(val).(*string); ok {
		if val == nil || len(*casted) == 0 {
			return nil
		}
		return val
	}
	if casted, ok := any(val).(*time.Time); ok {
		empty := time.Time{}
		if val == nil || (*casted).IsZero() || *casted == empty {
			return nil
		}
		return val
	}
	return nil
}

func TrimEmptyOrNil(val *string) *string {
	*val = strings.TrimSpace(*val)
	if val == nil || len(*val) == 0 {
		return nil
	}
	return val
}
