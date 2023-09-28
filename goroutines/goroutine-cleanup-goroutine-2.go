package main

import (
	"fmt"
)

func countTo2(max int, done chan struct{}) <-chan int {
	ch := make(chan int)
	i := 0
	go func() {
		for {
			select {
			case ch <- i:
				i++
			case <-done:
				fmt.Println("goroutine done")
			}
		}
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	result := countTo2(10, done)
	for i := range result {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	// We are closing done explicitly
	close(done)
}
