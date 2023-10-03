package main

import (
	"fmt"
	"sync"
)

func main() {
	numGoroutines := 5
	processorFunc2 := processor2
	input := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			input <- i
		}
		// Make sure to close the channel
		// For-range will read from this
		close(input)
	}()
	fmt.Println(processAndGather2(input, processorFunc2, numGoroutines))

}
func processAndGather2(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)

	var wg sync.WaitGroup
	wg.Add(num)

	// Spawn go routines and process the input concurrently
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	// Wait for all goroutines to finish
	// Also close the channel used for writing
	// That will also be used by for-range for reading
	go func() {
		wg.Wait()
		close(out)
	}()

	// Gather the data from the channel to slice
	var result []int
	for v := range out {
		result = append(result, v)
	}

	// return result slice
	return result
}

func processor2(i int) int {
	return i * 2
}
