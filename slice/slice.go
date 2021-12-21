package slice

import (
	"constraints"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Each calls the function on each item in the slice.
func Each[A ~[]T, T any](arr A, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

// Collect returns a new slice of values by mapping each value of original slice through a transformation function.
func Collect[A ~[]T, T any, M any](arr A, f func(T) M) []M {
	ret := make([]M, len(arr))

	for i, v := range arr {
		ret[i] = f(v)
	}

	return ret
}

// Reduce reduces a slice of values to single value.
func Reduce[A ~[]T, T any, M any](arr A, f func(M, T) M, initial M) M {
	for _, v := range arr {
		initial = f(initial, v)
	}

	return initial
}

// Find returns the first element in the slice that matches the condition.
// If slice doesn't contain an element it returns a default type value and false as second value.
func Find[A ~[]T, T any](arr A, f func(T) bool) (T, bool) {
	for _, v := range arr {
		if f(v) {
			return v, true
		}
	}

	return defaultvalue[T](), false
}

// Filter returns all elements in the slice that mathch the condition.
func Filter[A ~[]T, T any](arr A, f func(T) bool) A {
	var ret A

	for _, v := range arr {
		if f(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

// Every returns true if all elements match the condition.
func Every[A ~[]T, T any](arr A, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}

	return true
}

// Some returns true if there is at least one element that satisfies the condition.
func Some[A ~[]T, T any](arr A, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}

	return false
}

// Contains returns true if value is present in the slice.
func Contains[A ~[]T, T comparable](arr A, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

// Max returns the maximum value from the slice.
// If input slice is empty it returns a default value for input type.
func Max[A ~[]T, T constraints.Ordered](arr A) T {
	if len(arr) == 0 {
		return defaultvalue[T]()
	}

	e := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > e {
			e = arr[i]
		}
	}

	return e
}

// Min returns the minimum value from the slice.
// If input slice is empty it returns a default value for input type.
func Min[A ~[]T, T constraints.Ordered](arr A) T {
	if len(arr) == 0 {
		return defaultvalue[T]()
	}

	e := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < e {
			e = arr[i]
		}
	}

	return e
}

// GroupBy splits the slice into groups, grouped by the result of the function call.
func GroupBy[A ~[]T, T any, M comparable](arr A, f func(T) M) map[M]A {
	ret := make(map[M]A)

	for _, v := range arr {
		m := f(v)
		ret[m] = append(ret[m], v)
	}

	return ret
}

// Sample returns the random element from slice.
func Sample[A ~[]T, T any](arr A) T {
	return arr[rand.Intn(len(arr))]
}

// SampleN returns the N random elements from slice.
func SampleN[A ~[]T, T any](arr A, n int) []T {
	if n < 0 {
		return A{}
	}

	if n > len(arr) {
		n = len(arr)
	}

	ret := make([]T, n)
	for i, v := range rand.Perm(n) {
		ret[i] = arr[v]
	}

	return ret
}

// Union returns a slice of unique values from passed slices.
func Union[A ~[]T, T comparable](arr ...A) A {
	if len(arr) == 0 {
		return A{}
	}

	if len(arr) == 1 {
		return arr[0]
	}

	ret := make(A, 0, len(arr[0]))
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

// Intersection returns a slice of values that are in all passed slices.
func Intersection[A ~[]T, T comparable](arr ...A) A {
	if len(arr) == 0 {
		return A{}
	}

	if len(arr) == 1 {
		return arr[0]
	}

	ret := arr[0]
	arr = arr[1:]

	for len(arr) != 0 {
		var nextPath A

		part2 := arr[0]
		m := make(map[T]struct{})

		for _, array := range []A{ret, part2} {
			for i := 0; i < len(array); i++ {
				if _, ok := m[array[i]]; ok {
					nextPath = append(nextPath, array[i])
				} else {
					m[array[i]] = struct{}{}
				}
			}
		}

		ret = nextPath
		arr = arr[1:]
	}

	return ret
}

// Uniq returns a slice of unique values.
func Uniq[A ~[]T, T comparable](arr A) []T {
	ret := make(A, 0)
	m := make(map[T]struct{})

	for _, elem := range arr {
		if _, ok := m[elem]; !ok {
			m[elem] = struct{}{}
			ret = append(ret, elem)
		}
	}

	return ret
}

// IndexOf returns first index of the found element in the slice.
// If slice doesn't contain an element it returns -1.
func IndexOf[A ~[]T, T comparable](arr A, value T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}

	return -1
}

// LastIndexOf like as IndexOf, but the search goes from the end.
func LastIndexOf[A ~[]T, T comparable](arr A, value T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == value {
			return i
		}
	}

	return -1
}

// Reverse reverses the order of the elements in place.
func Reverse[A ~[]T, T any](arr A) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

func defaultvalue[T any]() T {
	n := new(T)
	return *n
}