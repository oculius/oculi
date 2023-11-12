package maputils

func ToArray[C comparable, V any](data map[C]V) []V {
	result := make([]V, len(data))
	i := 0
	for _, each := range data {
		result[i] = each
		i++
	}
	return result
}
