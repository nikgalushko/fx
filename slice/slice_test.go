package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEach(t *testing.T) {
	sum := 0
	Each([]int{1, 2, 3, 4}, func(i int) {
		sum += i
	})

	require.Equal(t, 10, sum)
}

func TestCollect(t *testing.T) {
	arr := Collect([]int{1, 2, 3, 4}, func(i int) int {
		return i * 2
	})

	require.Equal(t, []int{2, 4, 6, 8}, arr)
}

func TestReduce(t *testing.T) {
	join := Reduce([]string{"b", "l", "a", "h"}, func(memo, s string) string {
		return memo + s
	}, "")

	require.Equal(t, "blah", join)
}

func TestFind(t *testing.T) {
	tests := map[string]struct {
		arr     []int
		f       func(int) bool
		element int
		ok      bool
	}{
		"found": {
			arr:     []int{1, 2, 3, 4, 5},
			f:       func(i int) bool { return i == 4 },
			element: 4,
			ok:      true,
		},
		"not found": {
			arr:     []int{1, 2, 3, 4, 5},
			f:       func(i int) bool { return i == 100 },
			element: 0,
			ok:      false,
		},
	}

	for title, tt := range tests {
		element, ok := Find(tt.arr, tt.f)

		require.Equal(t, tt.element, element, title)
		require.Equal(t, tt.ok, ok, title)
	}
}

func TestFilter(t *testing.T) {
	ret := Filter([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i < 10 })

	require.Equal(t, []int{1, 4, 5, 2}, ret)
}

func TestEvery(t *testing.T) {
	tests := map[string]struct {
		arr      []int
		f        func(int) bool
		expected bool
	}{
		"not ok": {
			arr:      []int{10, 1, 4, 20, 5, 2},
			f:        func(i int) bool { return i >= 10 },
			expected: false,
		},
		"ok": {
			arr:      []int{10, 1, 4, 20, 5, 2},
			f:        func(i int) bool { return i >= 0 },
			expected: true,
		},
	}

	for title, tt := range tests {
		require.Equal(t, tt.expected, Every(tt.arr, tt.f), title)
	}
}

func TestSome(t *testing.T) {
	tests := map[string]struct {
		arr      []int
		f        func(int) bool
		expected bool
	}{
		"not ok": {
			arr:      []int{10, 1, 4, 20, 5, 2},
			f:        func(i int) bool { return i < 0 },
			expected: false,
		},
		"ok": {
			arr:      []int{10, 1, 4, 20, 5, 2},
			f:        func(i int) bool { return i >= 10 },
			expected: true,
		},
	}

	for title, tt := range tests {
		require.Equal(t, tt.expected, Some(tt.arr, tt.f), title)
	}
}

func TestContains(t *testing.T) {
	tests := map[string]struct {
		arr      []int
		value    int
		expected bool
	}{
		"contains": {
			arr:      []int{1, 2, 10, 23, 4},
			value:    4,
			expected: true,
		},
		"doesn't contain": {
			arr:      []int{1, 2, 10, 23, 4},
			value:    423,
			expected: false,
		},
	}

	for title, tt := range tests {
		require.Equal(t, tt.expected, Contains(tt.arr, tt.value), title)
	}
}

func TestGroupBy(t *testing.T) {
	group := GroupBy([]string{"one", "two", "three"}, func(s string) int { return len(s) })

	require.Equal(t, map[int][]string{3: {"one", "two"}, 5: {"three"}}, group)
}

func TestSample(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	v := Sample(arr)

	require.Equal(t, true, Contains(arr, v))
}

func TestSampleN(t *testing.T) {
	check := func(arr, samples []int){
		require.Equal(t, len(Uniq(samples)), len(samples))

		for _, v := range samples {
			require.Equal(t, true, Contains(arr, v))
		}
	}

	tests := []struct{
		input []int
		n int
	}{
		{
			input: []int{11, 12, 13, 14, 15, 16, 17, 18, 19},
			n: 5,
		},
		{
			input: []int{0, 1, 2},
			n: 1,
		},
		{
			input: []int{0, 1, 2},
			n: 100,
		},
		{
			input: []int{0, 1, 2},
			n: -10,
		},
	}
	
	for _, tt := range tests {
		check(tt.input, SampleN(tt.input, tt.n))
	}
}

func TestUnion(t *testing.T) {
	tests := map[string]struct {
		in       [][]string
		expected []string
	}{
		"two arrays": {
			in:       [][]string{{"a", "b", "c"}, {"b", "c", "d"}},
			expected: []string{"a", "b", "c", "d"},
		},
		"zero arrays": {
			in: nil,
			expected: []string{},
		},
		"one array": {
			in:       [][]string{{"1", "2", "3"}},
			expected: []string{"1", "2", "3"},
		},
		"three arrays": {
			in:       [][]string{{"a", "b", "c"}, {"b", "c", "d"}, {"c", "d", "e"}},
			expected: []string{"a", "b", "c", "d", "e"},
		},
	}

	for title, tt := range tests {
		require.Equal(t, tt.expected, Union(tt.in...), title)
	}
}

func TestIntersection(t *testing.T) {
	tests := map[string]struct{
		in [][]string
		expected []string
	}{
		"without arrays": {
			in: [][]string{},
			expected: []string{},
		},
		"single array": {
			in: [][]string{{"a", "a", "b"}},
			expected: []string{"a", "a", "b"},
		},
		"two arrays": {
			in: [][]string{{"a", "b", "c"}, {"b", "c", "d"}},
			expected: []string{"b", "c"},
		},
		"three arrays": {
			in: [][]string{{"a", "b", "c"}, {"b", "c", "d"}, {"a", "f", "g", "c"}},
			expected: []string{"c"},
		},
		"four arrays": {
			in: [][]string{{"a", "b", "c"}, {"b", "c", "d"}, {"a", "f", "g", "c"}, {"k", "v"}},
			expected: nil,
		},
	}

	for title, tt := range tests {
		require.Equal(t, tt.expected, Intersection(tt.in...), title)
	}
}

func TestIndexOf(t *testing.T) {
	require.Equal(t, 1, IndexOf([]int{1, 2, 3, 2}, 2), "found")
	require.Equal(t, -1, IndexOf([]int{1, 2, 3, 4}, 20), "not found")
}

func TestLastIndexOf(t *testing.T) {
	require.Equal(t, 3, LastIndexOf([]int{1, 2, 3, 2}, 2), "found")
	require.Equal(t, -1, LastIndexOf([]int{1, 2, 3, 4}, 20), "not found")
}

func TestMinMax(t *testing.T) {
	tests := map[string]struct{
		in []int
		max int
		min int
	}{
		"simple array": {
			in: []int{10, 2, 1, 4, 19},
			max: 19,
			min: 1,
		},
		"zero array": {
			in : nil,
			max: 0,
			min: 0,
		},
	}
	
	for title, tt := range tests {
		require.Equal(t, tt.max, Max(tt.in), title)
		require.Equal(t, tt.min, Min(tt.in), title)
	}
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
