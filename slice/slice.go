package main


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