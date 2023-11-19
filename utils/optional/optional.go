package optional

import (
	"strconv"
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

func Zero[T Number | time.Duration](val *T) *T {
	if val == nil || *val == T(0) {
		return nil
	}
	return val
}

var emptyTime = time.Time{}

func Empty[T string | time.Time](val *T) *T {
	if casted, ok := any(val).(*string); ok {
		if val == nil || len(*casted) == 0 {
			return nil
		}
		return val
	}
	if casted, ok := any(val).(*time.Time); ok {
		if val == nil || (*casted).IsZero() || *casted == emptyTime {
			return nil
		}
		return val
	}
	return nil
}

func Trim(val *string) *string {
	if val == nil {
		return nil
	}
	*val = strings.TrimSpace(*val)
	if len(*val) == 0 {
		return nil
	}
	return val
}

func Bool(val string, def bool) bool {
	res, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return res
}
