package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} { // 1

		// Base case for recursion
		switch len(channels) {
		case 0: // 2 - No of channels is zero
			return nil
		case 1: // 3 - No of channels is one
			return channels[0]
		}

		orDone := make(chan interface{})

		go func() { // 4 Main recursive body
			defer close(orDone)

			switch len(channels) {

			case 2: // 5 - No of channels is two
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default: // 6 - For all other counts of channels
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // 6 - Recursive call
				}
			}
		}()

		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))

}
