package slice

import (
	"constraints"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Each[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func Collect[T any, M any](arr []T, f func(T) M) []M {
	ret := make([]M, len(arr))

	for i, v := range arr {
		ret[i] = f(v)
	}

	return ret
}

func Reduce[T any](arr []T, f func(T, T) T, initial T) T {
	for _, v := range arr {
		initial = f(initial, v)
	}

	return initial
}

func Find[T any](arr []T, f func(T) bool) (T, bool) {
	for _, v := range arr {
		if f(v) {
			return v, true
		}
	}

	n := new(T)
	return *n, false
}

func Filter[T any](arr []T, f func(T) bool) []T {
	var ret []T

	for _, v := range arr {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func Every[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}

	return true
}

func Some[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}

	return false
}

func Contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func Max[T constraints.Ordered](arr []T) T {
	e := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > e {
			e = arr[i]
		}
	}

	return e
}

func Min[T constraints.Ordered](arr []T) T {
	e := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < e {
			e = arr[i]
		}
	}

	return e
}

func GroupBy[T any, M comparable](arr []T, f func(T) M) map[M][]T {
	ret := make(map[M][]T)

	for _, v := range arr {
		m := f(v)
		ret[m] = append(ret[m], v)
	}

	return ret
}

func Sample[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}

func SampleN[T any](arr []T, n int) []T {
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

func Union[T comparable](arr ...[]T) []T {
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

func Intersection[T comparable](arr ...[]T) []T {
	m := make(map[T]struct{})
	var ret []T

	for _, array := range arr {
		for i := 0; i < len(array); i++ {
			if _, ok := m[array[i]]; ok {
				ret = append(ret, array[i])
			}
			m[array[i]] = struct{}{}
		}
	}

	return ret
}

func Uniq[T comparable](arr []T) []T {
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

func IndexOf[T comparable](arr []T, value T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}

	return -1
}

func LastIndexOf[T comparable](arr []T, value T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == value {
			return i
		}
	}

	return -1
}
