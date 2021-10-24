package slice

import (
	"fmt"
	"testing"
	"reflect"
)

func TestEach(t *testing.T) {
	sum := 0
	each([]int{1, 2, 3, 4}, func(i int) {
		sum += i
	})

	equal(10, sum, "")
}

func TestCollect(t *testing.T) {
	arr := collect([]int{1, 2, 3, 4}, func(i int) int {
		return i * 2
	})

	equal([]int{2, 4, 6, 8}, arr, "TestCollect")
}

func TestReduct(t *testing.T) {
	join := reduce([]string{"b", "l", "a", "h"}, func(memo, s string) string {
		return memo + s
	}, "")

	equal("blah", join, "TestReduct")
}

func TestFind(t *testing.T) {
	element, ok := find([]int{1, 2, 3, 4, 5}, func(i int) bool { return i == 4 })

	equal(4, element, "TestFind")
	equal(true, ok, "TestFind")

	element, ok = find([]int{1, 2, 3, 4, 5}, func(i int) bool { return i == 100 })

	equal(0, element, "TestFind")
	equal(false, ok, "TestFind")
}

func TestFilter(t *testing.T) {
	ret := filter([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i < 10 })

	equal(ret, []int{1, 4, 5, 2}, "TestFilter")
}

func TestEvery(t *testing.T) {
	ret := every([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i >= 10 })

	equal(ret, false, "TestEvery")

	ret = every([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i >= 0 })

	equal(ret, true, "TestEvery")
}

func TestSome(t *testing.T) {
	ret := some([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i >= 10 })

	equal(ret, true, "TestSome")

	ret = every([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i < 0 })

	equal(ret, false, "TestSome")
}

func TestGroupBy(t *testing.T) {
	group := groupBy([]string{"one", "two", "three"}, func(s string) int { return len(s) })

	equal(group, map[int][]string{3: {"one", "two"}, 5: {"three"}}, "TestGroupBy")
}

func TestSample(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	v := sample(arr)

	equal(true, contains(arr, v), "TestSample")
}

func TestSampleN(t *testing.T) {
	arr := []int{11, 12, 13, 14, 15, 16, 17, 18, 19}
	samples := sampleN(arr, 5)

	equal(len(uniq(samples)), len(samples), "TestSampleN")
	for _, v := range samples {
		equal(true, contains(arr, v), "TestSampleN")
	}
}

func TestUnion(t *testing.T) {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}

	equal(union(arr1, arr2), []string{"a", "b", "c", "d"}, "TestUnion")
}

func TestIntersection(t *testing.T) {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}

	equal(intersection(arr1, arr2), []string{"b", "c"}, "TestIntersaction")
}

func TestIndexOf(t *testing.T) {
	equal(indexOf([]int{1, 2, 3, 2}, 2), 1, "TestIndexOf")
	equal(indexOf([]int{1, 2, 3, 4}, 20), -1, "TestIndexOf not found")
}

func TestLastIndexOf(t *testing.T) {
	equal(lastIndexOf([]int{1, 2, 3, 2}, 2), 3, "TestLastIndexOf")
	equal(lastIndexOf([]int{1, 2, 3, 4}, 20), -1, "TestLastIndexOf not found")
}

func TestMinMax(t *testing.T) {
	arr := []int{10, 2, 1,4, 19}
	
	equal(max(arr), 19, "TestMinMax")
	equal(min(arr), 1, "TestMinMax")
}

func equal[T any](actual, expected T, title string) {
	if !reflect.DeepEqual(actual, expected) {
		panic(fmt.Sprintf("actual %v != expected %v", actual, expected))
	} else {
		fmt.Println(title + " - OK")
	}
}