package kv

// Keys reurns a slice of keys from map.
func Keys[K comparable, V any](m map[K]V) []K {
	ret := make([]K, 0, len(m))

	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

// Values reurns a slice of values from map.
func Values[K comparable, V any](m map[K]V) []V {
	ret := make([]V, 0, len(m))

	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

// Each calls the function on each key-value pair of map.
func Each[K comparable, V any](m map[K]V, f func(key K, value V)) {
	for k, v := range m {
		f(k, v)
	}
}

// Filter returns a new map that contains key-value pairs that mathch the condition.
func Filter[K comparable, V any](m map[K]V, f func(key K, value V) bool) map[K]V {
	ret := make(map[K]V)

	for k, v := range m {
		if f(k, v) {
			ret[k] = v
		}
	}

	return ret
}
