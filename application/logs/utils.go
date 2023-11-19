package logs

import "fmt"

func anyToString(item any) string {
	key, ok := item.(string)
	if ok {
		return key
	}

	stringer, ok := item.(fmt.Stringer)
	if ok {
		return stringer.String()
	}

	return fmt.Sprintf("%#v", item)
}
