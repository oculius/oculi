package arrayutils

import "github.com/oculius/oculi/v2/utils/maputils"

func Identity[V any](v V) V {
	return v
}

func ToMap[C comparable, V any](array []V, identifier func(V) C) map[C][]V {
	result := map[C][]V{}
	for _, each := range array {
		group := identifier(each)
		result[group] = append(result[group], each)
	}
	return result
}

func ToMapUnique[C comparable, V any](data []V, identifier func(V) C) (unique map[C]V, omitted int) {
	result := map[C]V{}
	omitted = 0
	for _, each := range data {
		id := identifier(each)
		if _, ok := result[id]; !ok {
			result[id] = each
		} else {
			omitted++
		}
	}
	unique = result
	return
}

func Unique[C comparable, V any](data []V, identifier func(V) C) (unique []V, omitted int) {
	uniqueMap, omittedVal := ToMapUnique[C, V](data, identifier)
	unique = maputils.ToArray[C, V](uniqueMap)
	omitted = omittedVal
	return
}
