package main

import (
	"fmt"
	"time"
)

func main() {
	result, done := multiples(6)
	for i := range result {
		if i > 20 {
			break
		}
		fmt.Println(i)
	}
	close(done)
	time.Sleep(2 * time.Second)
}

// multiples return a bi-directional channel for closing
func multiples(i int) (<-chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	current := 0
	go func() {
		for {
			select {
			case out <- current * i:
				current++
			case <-done:
				fmt.Println("goroutine done")
				return
			}
		}
	}()
	return out, done
}
