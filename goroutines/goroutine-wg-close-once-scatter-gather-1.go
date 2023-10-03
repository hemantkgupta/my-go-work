package main

import (
	"fmt"
	"sync"
)

func main() {
	numGoroutines := 5
	/*
		processorFunc1 := func(i int) int {
			return i * 2
		}
	*/
	processorFunc2 := processor
	fmt.Println(processAndGather(numGoroutines, processorFunc2))

}

func processAndGather(num int, myProcessor func(int) int) []int {
	fmt.Println("Processing data.")
	out := make(chan int, num)

	var wg sync.WaitGroup
	wg.Add(num)

	// Spawn go routines and process the input concurrently
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("From goroutine", i)
			out <- myProcessor(i)
		}(i)
	}

	wg.Wait()
	fmt.Println("The wait completed.")
	close(out)
	var result []int
	for v := range out {
		result = append(result, v)
	}
	// return result slice
	return result
}

func processor(i int) int {
	return i * 2
}
