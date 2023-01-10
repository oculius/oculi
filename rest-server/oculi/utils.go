package oculi

import (
	"reflect"
	"runtime"
	_ "unsafe"
)

func handlerName(h any) string {
	t := reflect.ValueOf(h).Type()
	if t.Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
	}
	return t.String()
}
