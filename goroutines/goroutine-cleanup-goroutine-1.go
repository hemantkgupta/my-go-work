package main

import (
	"fmt"
)

func countTo(max int) (<-chan int, chan struct{}) {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		for i := 0; i < max; i++ {
			select {
			case ch <- i:
			case <-done:
				fmt.Println("goroutine done")
			}
		}
	}()
	return ch, done
}

func main() {
	result, done := countTo(10)
	for i := range result {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	close(done)
}
