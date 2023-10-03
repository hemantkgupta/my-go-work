package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doThing1()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
	}()
	wg.Wait()
}

func doThing1() {
	time.Sleep(1 * time.Second)
}

func doThing2() {
	time.Sleep(2 * time.Second)
}

func doThing3() {
	time.Sleep(3 * time.Second)
}
