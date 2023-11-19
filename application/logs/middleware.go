package logs

import "fmt"

type (
	Middleware func(info map[string]any)
)

func maskAll(input string) string {
	masked := ""
	for range input {
		masked += "*"
	}
	return masked
}

func maskString(input string, visibleStart, visibleEnd int) string {
	if visibleStart+visibleEnd >= len(input) {
		return maskAll(input)
	}

	visiblePart := input[:visibleStart] + input[len(input)-visibleEnd:]
	maskedPart := maskAll(input[visibleStart : len(input)-visibleEnd])

	return visiblePart + maskedPart
}

func run(i map[string]any, applierFn func(string) string, keys []string) {
	if len(i) == 0 {
		return
	}
	for _, key := range keys {
		val, ok := i[key]
		if !ok || val == nil {
			continue
		}
		switch v := val.(type) {
		case map[string]any:
			run(v, applierFn, keys)
		case string:
			i[key] = applierFn(v)
		case fmt.Stringer:
			i[key] = applierFn(v.String())
		default:
			i[key] = applierFn(anyToString(v))
		}
	}
}

func NewApplierMiddleware(applierFn func(string) string, keys ...string) Middleware {
	if len(keys) == 0 {
		panic("No key provided while creating a new mask middleware")
	}
	return func(info map[string]any) {
		run(info, applierFn, keys)
	}
}

func NewMaskMiddleware(visibleStart, visibleEnd int, keys ...string) Middleware {
	return NewApplierMiddleware(func(input string) string {
		return maskString(input, visibleStart, visibleEnd)
	}, keys...)
}
