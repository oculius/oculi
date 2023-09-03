package utils

func ArrayToMap[C comparable, V any](data []V, identifier func(V) C) map[C][]V {
	result := map[C][]V{}
	for _, each := range data {
		group := identifier(each)
		result[group] = append(result[group], each)
	}
	return result
}

func ArrayToMapUnique[C comparable, V any](data []V, identifier func(V) C) (unique map[C]V, omitted int) {
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

func MapToArray[C comparable, V any](data map[C]V) []V {
	result := make([]V, len(data))
	i := 0
	for _, each := range data {
		result[i] = each
		i++
	}
	return result
}

func ArrayUnique[C comparable, V any](data []V, identifier func(V) C) (unique []V, omitted int) {
	uniqueMap, omittedVal := ArrayToMapUnique[C, V](data, identifier)
	unique = MapToArray[C, V](uniqueMap)
	omitted = omittedVal
	return
}
