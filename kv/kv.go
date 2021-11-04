package kv

func Keys[K comparable, V any](m map[K]V) []K {
	ret := make([]K, 0, len(m))

	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

func Values[K comparable, V any](m map[K]V) []V {
	ret := make([]V, 0, len(m))

	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

func Each[K comparable, V any](m map[K]V, f func(key K, value V)) {
	for k, v := range m {
		f(k, v)
	}
}
