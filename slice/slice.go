package slice

import (
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func each[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func collect[T any, M any](arr []T, f func(T) M) []M {
	ret := make([]M, len(arr))

	for i, v := range arr {
		ret[i] = f(v)
	}

	return ret
}

func reduce[T any](arr []T, f func(T,T) T, initial T) T {
	for _, v := range arr {
		initial = f(initial, v)
	}

	return initial
}

func find[T any](arr []T, f func(T) bool) (T, bool) {
	for _, v := range arr {
		if f(v) {
			return v, true
		}
	}

	n := new(T)
	return *n, false
}

func filter[T any](arr []T, f func(T) bool) []T {
	var ret []T

	for _, v := range arr {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func every[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}

	return true
}

func some[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}

	return false
}

func contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func max() {}
func min() {}

func groupBy[T any, M comparable](arr []T, f func(T) M) map[M][]T {
	ret := make(map[M][]T)

	for _, v := range arr {
		m := f(v)
		ret[m] = append(ret[m], v)
	}

	return ret
}

func sample[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}

func sampleN[T any](arr []T, n int) []T {
	indexes := make([]int, len(arr))
	for i := 0; i < len(indexes); i++ {
		indexes[i] = i
	}

	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	ret := make([]T, n)
	for i := 0; i < n; i++ {
		ret[i] = arr[indexes[i]]
	}

	return ret
}

func union[T comparable](arr ...[]T) []T {
	if len(arr) == 0 {
		return nil
	}

	if len(arr) == 1 {
		return arr[0]
	}

	ret := make([]T, 0, len(arr[0]))
	m := make(map[T]struct{})
	for _, array := range arr {
		for i := 0; i < len(array); i++ {
			if _, ok := m[array[i]]; !ok {
				ret = append(ret, array[i])
				m[array[i]] = struct{}{}
			}
		}
	}

	return ret
}

func intersection[T comparable](arr ...[]T) []T {
	m := make(map[T]struct{})
	for _, array := range arr {
		for i := 0; i < len(array); i++ {
			m[array[i]] = struct{}{}
		}
	}

	ret := make([]T, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

func uniq[T comparable](arr []T) []T {
	m := make(map[T]struct{})
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = struct{}{}
	}


	ret := make([]T, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

func indexOf[T comparable](arr []T, value T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}

	return -1
}

func lastIndexOf[T comparable](arr []T, value T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == value {
			return i
		}
	}

	return -1
}