package kv

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	keys := Keys(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	sort.Strings(keys)

	require.Equal(t, []string{"one", "three", "two"}, keys)
}

func TestValues(t *testing.T) {
	values := Values(map[int]int{
		0: 10,
		1: 11,
		2: 12,
	})

	sort.Ints(values)

	require.Equal(t, []int{10, 11, 12}, values)
}

func TestEach(t *testing.T) {
	var (
		keys   []int
		values []string
	)

	Each(map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}, func(key int, value string) {
		if key > 3 {
			return
		}

		keys = append(keys, key)
		values = append(values, value)
	})

	sort.Ints(keys)
	sort.Strings(values)

	require.Equal(t, []int{1, 2, 3}, keys)
	require.Equal(t, []string{"a", "b", "c"}, values)
}

func TestFilter(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		5: "five",
		6: "six",
		7: "seven",
	}

	Filter(m, func(key int, value string) bool {
		return key < 3 || len(value) < 4
	})

	require.Equal(t, map[int]string{5: "five", 7: "seven"}, m)
}
