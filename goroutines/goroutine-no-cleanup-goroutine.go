package main

import "fmt"

func countTo1(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {

	noCleanUp()

}

func withCleanUp() {
	// Used for-range on channel
	for i := range countTo1(10) {
		fmt.Println(i)
	}
}

func noCleanUp() {
	for i := range countTo1(10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
