package ch

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/nikgalushko/fx/slice"
	"github.com/stretchr/testify/require"
)

func TestFunnel(t *testing.T) {
	var channels []chan int
	for i := 0; i < 3; i++ {
		var arr []int
		for j := 0; j < 5; j++ {
			arr = append(arr, j)
		}

		channels = append(channels, createChanFrom(arr))
	}

	ch := Funnel(channels...)
	sum := 0

	for v := range ch {
		sum += v
	}

	if sum != 30 {
		t.Fatalf("%d != 30", sum)
	}
}

func TestSplit(t *testing.T) {
	ch := createChanFrom([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	splitted := Split(ch, 3)
	require.Equal(t, 3, len(splitted))

	var wg sync.WaitGroup
	wg.Add(3)

	sum := new(int32)
	slice.Each(splitted, func(c <-chan int) {
		for v := range c {
			atomic.AddInt32(sum, int32(v))
		}
		wg.Done()
	})

	wg.Wait()
	require.Equal(t, int32(45), *sum)
}

func TestSplitPanic(t *testing.T) {
	ch := createChanFrom([]int{1, 2})

	require.Panics(t, func() { Split(ch, 0) })
	require.Panics(t, func() { Split(ch, -1) })
}

func createChanFrom(arr []int) chan int {
	ret := make(chan int, len(arr))
	for _, v := range arr {
		ret <- v
	}
	close(ret)

	return ret
}
