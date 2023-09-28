package main

import "fmt"

func countTo3(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancelFunc := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:

			}
		}
		close(ch)
	}()
	return ch, cancelFunc
}

func main() {
	ch, cancel := countTo3(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	// Call the cancel to close done channel
	cancel()
}
