package slice

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEach(t *testing.T) {
	sum := 0
	Each([]int{1, 2, 3, 4}, func(i int) {
		sum += i
	})

	equal(10, sum, "")
}

func TestCollect(t *testing.T) {
	arr := Collect([]int{1, 2, 3, 4}, func(i int) int {
		return i * 2
	})

	equal([]int{2, 4, 6, 8}, arr, "TestCollect")
}

func TestReduct(t *testing.T) {
	join := Reduce([]string{"b", "l", "a", "h"}, func(memo, s string) string {
		return memo + s
	}, "")

	equal("blah", join, "TestReduct")
}

func TestFind(t *testing.T) {
	element, ok := Find([]int{1, 2, 3, 4, 5}, func(i int) bool { return i == 4 })

	equal(4, element, "TestFind")
	equal(true, ok, "TestFind")

	element, ok = Find([]int{1, 2, 3, 4, 5}, func(i int) bool { return i == 100 })

	equal(0, element, "TestFind")
	equal(false, ok, "TestFind")
}

func TestFilter(t *testing.T) {
	ret := Filter([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i < 10 })

	equal(ret, []int{1, 4, 5, 2}, "TestFilter")
}

func TestEvery(t *testing.T) {
	ret := Every([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i >= 10 })

	equal(ret, false, "TestEvery")

	ret = Every([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i >= 0 })

	equal(ret, true, "TestEvery")
}

func TestSome(t *testing.T) {
	ret := Some([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i >= 10 })

	equal(ret, true, "TestSome")

	ret = Every([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i < 0 })

	equal(ret, false, "TestSome")
}

func TestGroupBy(t *testing.T) {
	group := GroupBy([]string{"one", "two", "three"}, func(s string) int { return len(s) })

	equal(group, map[int][]string{3: {"one", "two"}, 5: {"three"}}, "TestGroupBy")
}

func TestSample(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	v := Sample(arr)

	equal(true, Contains(arr, v), "TestSample")
}

func TestSampleN(t *testing.T) {
	arr := []int{11, 12, 13, 14, 15, 16, 17, 18, 19}
	samples := SampleN(arr, 5)

	equal(len(Uniq(samples)), len(samples), "TestSampleN")
	for _, v := range samples {
		equal(true, Contains(arr, v), "TestSampleN")
	}
}

func TestUnion(t *testing.T) {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}

	equal(Union(arr1, arr2), []string{"a", "b", "c", "d"}, "TestUnion")
}

func TestIntersection(t *testing.T) {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}

	equal(Intersection(arr1, arr2), []string{"b", "c"}, "TestIntersaction")
}

func TestIndexOf(t *testing.T) {
	equal(IndexOf([]int{1, 2, 3, 2}, 2), 1, "TestIndexOf")
	equal(IndexOf([]int{1, 2, 3, 4}, 20), -1, "TestIndexOf not found")
}

func TestLastIndexOf(t *testing.T) {
	equal(LastIndexOf([]int{1, 2, 3, 2}, 2), 3, "TestLastIndexOf")
	equal(LastIndexOf([]int{1, 2, 3, 4}, 20), -1, "TestLastIndexOf not found")
}

func TestMinMax(t *testing.T) {
	arr := []int{10, 2, 1, 4, 19}

	equal(Max(arr), 19, "TestMinMax")
	equal(Min(arr), 1, "TestMinMax")
}

func TestReverse(t *testing.T) {
	tests := map[string]struct {
		in       []int
		expected []int
	}{
		"even": {
			in:       []int{0, 1, 2, 3},
			expected: []int{3, 2, 1, 0},
		},
		"odd": {
			in:       []int{0, 1, 2, 3, 4},
			expected: []int{4, 3, 2, 1, 0},
		},
	}

	for title, tt := range tests {
		Reverse(tt.in)
		require.Equal(t, tt.expected, tt.in, title)
	}
}

func equal[T any](actual, expected T, title string) {
	if !reflect.DeepEqual(actual, expected) {
		panic(fmt.Sprintf("actual %v != expected %v", actual, expected))
	} else {
		fmt.Println(title + " - OK")
	}
}
