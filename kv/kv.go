package kv

import (
	"github.com/nikgalushko/fx/slice"
)

func Keys[K any, V any](m map[K]V) []K {
	ret := make([]K, 0, len(m))

	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

func Values[K any, V any](m map[K]V) []V {
	ret := make([]V, 0, len(m))

	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

func Each[K any, V any](m map[K]V, f (key K, value V)) {
	for k, v := range m {
		f(k, v)
	}
}

func Union(K any, V any)(maps ...map[K]V) map[K]V {
	maxSize := slice.Max(slice.Collect(maps, func(m map[K]V) int { len(m) }))
	ret := make(map[K]V, maxSize)

	for _, m := range maps {
		for k, v := range m {
			ret[k] = v
		}
	}

	return ret
}