package ch

import (
	"testing"
)

func TestMerge(t *testing.T) {
	var channels []chan int
	for i := 0; i < 3; i++ {
		var arr []int
		for j := 0; j < 5; j++ {
			arr = append(arr, j)
		}

		channels = append(channels, createChanFrom(arr))
	}

	ch := Merge(channels...)
	sum := 0

	for v := range ch {
		sum += v
	}

	if sum != 30 {
		t.Fatalf("%d != 30", sum)
	}
}

func createChanFrom(arr []int) chan int {
	ret := make(chan int, len(arr))
	for _, v := range arr {
		ret <- v
	}
	close(ret)

	return ret
}
