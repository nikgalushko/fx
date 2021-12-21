# Slice functions

## Each
Each calls the function on each item in the slice.

```
sum := 0
Each([]int{1, 2, 3, 4}, func(i int) {
  sum += i
})

// sum is 10
```
## Collect
Collect returns a new slice of values by mapping each value of original slice through a transformation function.

```
arr := Collect([]int{1, 2, 3, 4}, func(i int) int {
  return i * 2
})

// arr is []int{2, 4, 6, 8}
```
## Reduce
Reduce reduces a slice of values to single value.

```
join := Reduce([]string{"b", "l", "a", "h"}, func(memo, s string) string {
  return memo + s
}, "")

// join is "blah"
```

## Find
Find returns the first element in the slice that matches the condition. If slice doesn't contain an element it returns a default type value and false as second value.

```
element, ok := Find([]int{1, 2, 3, 4, 5}, func(i int) bool { return i == 4 })

// element is 4; ok is true
```

## Filter
Filter returns all elements in the slice that mathch the condition.

```
ret := Filter([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i < 10 })

// ret is []int{1, 4, 5, 2}
```

## Every
Every returns true if all elements match the condition.

```
ret := Every([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i >= 0 })

// ret is true
```

## Some
Some returns true if there is at least one element that satisfies the condition.

```
ret := Some([]int{10, 1, 4, 20, 5, 2}, func(i int) bool { return i < 0 })

// ret is false
```

## Contains
Contains returns true if value is present in the slice.

```
ret := Contains([]int{1, 2, 10, 23, 4}, 4)

// ret is true
```

## Max
Max returns the maximum value from the slice.
Is input slice is empty it returns a default value for input type.

```
arr := []int{10, 2, 1, 4, 19}
ret := Max(arr)

// ret is 19

zero := Max([]int{})

// zero is 0
```

## Min
Min returns the minimum value from the slice.
Is input slice is empty it returns a default value for input type.

```
arr := []int{10, 2, 1, 4, 19}
ret := Min(arr)

// ret is 11

zero := Max([]string{})

// zero is ""
```

## GroupBy
GroupBy splits the slice into groups, grouped by the result of the function call.

```
group := GroupBy([]string{"one", "two", "three"}, func(s string) int { return len(s) })

// group is map[int][]string{3: {"one", "two"}, 5: {"three"}}
```

## Sample
Sample returns the random element from slice.

```
v := Sample([]int{1, 2, 3, 4})

// v is random element from {1, 2, 3, 4}
```

## SampleN
SampleN returns the N random elements from slice.

```
arr := []int{11, 12, 13, 14, 15, 16, 17, 18, 19}
samples := SampleN(arr, 5)

// samples is slice of random elements from arr 
```

## Union
Union returns a slice of unique values from passed slices.

```
ret := Union([]string{"a", "b", "c"}, []string{"b", "c", "d"}}

// ret is []string{"a", "b", "c", "d"}
```

## Intersection
Intersection returns a slice of values that are in all passed slices.

```
arr1 := []string{"a", "b", "c"}
arr2 := []string{"b", "c", "d"}

ret := Intersection(arr1, arr2)

// ret is []string{"b", "c"}
```

## Uniq
Uniq returns a slice of unique values.

```
arr := []string{"a", "b", "a", "c", "b"}
ret := Uniq(arr)

// ret is []string{"a", "b", "c"}
```

## IndexOf
IndexOf returns first index of the found element in the slice. If slice doesn't contain an element it returns -1.

```
i := IndexOf([]int{1, 2, 3, 2}, 2)

// i is 1
```

## LastIndexOf
LastIndexOf like as IndexOf, but the search goes from the end.

```
i := LastIndexOf([]int{1, 2, 3, 4}, 20)

// i is -1
```

## Reduce
Reverse reverses the order of the elements in place.

```
arr := []int{0, 1, 2, 3}
Reverse(arr)

// arr is []int{3, 2, 1, 0}
```