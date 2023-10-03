package main

import "sync"

type Test struct {
	A int
}

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return &Test{
				A: 1,
			}
		},
	}

	testObject := pool.Get().(*Test)
	println(testObject.A) // print 1

	pool.Put(testObject)
}
