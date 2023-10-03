package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var a float64 = 0.0
var b float64 = 0.0

func done() {
	r := rand.New(rand.NewSource(99))
	lst1 := []int{}
	lst2 := []int{}
	lock1 := sync.Mutex{}
	lock2 := sync.RWMutex{}
	wtgrp1 := sync.WaitGroup{}
	wtgrp2 := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		lst1 = append(lst1, i)
		lst2 = append(lst2, i)
	}
	// 1. spawn 1000 goroutines and synchronize through locks
	now1 := time.Now()
	for i := 0; i < 1000; i++ {
		wtgrp1.Add(1)
		pos := r.Intn(10)
		go func() {
			defer wtgrp1.Done()
			lock1.Lock()
			defer lock1.Unlock()
			goLst := []int{} // goLst, destroy automatically
			// Process each element of lst1 randomly
			for j := 0; j < 1000; j++ {
				goLst = append(goLst, lst1[pos])
			}
		}()
	}
	wtgrp1.Wait()
	diff1 := time.Now().Sub(now1)
	// 2. spawn 1000 goroutines and synchronize through read-write locks
	now2 := time.Now()
	for i := 0; i < 1000; i++ {
		wtgrp2.Add(1)
		pos := r.Intn(10)
		go func() {
			defer wtgrp2.Done()
			lock2.RLock()
			defer lock2.RUnlock()
			goLst := []int{} // goLst, destroy automatically
			// Process each element of lst2 randomly
			for j := 0; j < 1000; j++ {
				goLst = append(goLst, lst1[pos])
			}
		}()
	}
	wtgrp2.Wait()
	diff2 := time.Now().Sub(now2)
	fmt.Println(diff1, "  ", diff2)
	a = a + float64(diff1)
	b = b + float64(diff2)
}
func main() {
	fmt.Println("Mutex   RWMutex")
	for i := 0; i < 10; i++ {
		done()
		time.Sleep(2 * time.Second)
	}
	fmt.Println(a, "  ", b)
	fmt.Println(b / a)
	fmt.Println("Faster by : ", ((a-b)*100)/a, " %")
}
