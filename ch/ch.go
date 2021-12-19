package ch

import (
	"sync"
)

func Merge[T any](channels ...chan T) chan T {
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
