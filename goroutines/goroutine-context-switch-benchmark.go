package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("test")

	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < 100; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < 100; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	//b.StartTimer()
	close(begin)
	wg.Wait()
}
