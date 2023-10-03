package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	val, err := timeLimit()
	if err != nil {
		fmt.Println("Error in processing: ", err)
	}
	fmt.Println(val)
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() (int, error) {
	//time.Sleep(1 * time.Second)
	time.Sleep(3 * time.Second)
	return 5, nil
}
