package ch

import (
	"sync"
)

// Funnel returns a channel containing the values from all channels.
func Funnel[T any](channels ...chan T) chan T {
	ret := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(c chan T) {
			for v := range c {
				ret <- v
			}

			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(ret)
	}()

	return ret
}

// Split implements a Fan-Out pattern.
func Split[T any](ch <-chan T, n int) []<-chan T {
	if n <= 0 {
		panic("splitting is only possible for a positive number of channels")
	}

	ret := make([]<-chan T, 0, n)

	for i := 0; i < n; i++ {
		c := make(chan T)

		go func() {
			defer close(c)
			for v := range ch {
				c <- v
			}
		}()

		ret = append(ret, c)
	}

	return ret
}
