# Map functions

## Keys
Keys reurns a slice of keys from map.

```
keys := Keys(map[string]int{
  "one":   1,
  "two":   2,
  "three": 3,
})


// keys is []string{"one", "three", "two"}. The order may be different.
```

## Values
Values reurns a slice of values from map.

```
values := Values(map[int]int{
  0: 10,
  1: 11,
  2: 12,
})

// values is []int{10, 11, 12}. The order may be different.
```

## Each
Each calls the function on each key-value pair of map.

```
m := map[int]string{
  1: "a",
  2: "b",
  3: "c",
  4: "d",
}

var (
  keys   []int
  values []string
)
Each(m, func(key int, value string) {
  if key > 3 {
    return
  }

  keys = append(keys, key)
  values = append(values, value)
})

// keys is []int{1, 2, 3}. The order may be different.
// values is []string{"a", "b", "c"}. The order may be different.
```

## Filter
Filter returns a new map that contains key-value pairs that mathch the condition.

```
m := map[int]string{
  1: "one",
  2: "two",
  5: "five",
  6: "six",
  7: "seven",
}

filtered := Filter(m, func(key int, value string) bool {
  return key < 3 || len(value) < 4
})

// filtered is map[int]string{1: "one", 2: "two", 6: "six"}. The order may be different. Original map is not changed.
```