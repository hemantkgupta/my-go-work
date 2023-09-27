package main

import "fmt"

func MyFunc1(opts MyFuncOpts1) {
	fmt.Println(opts.FirstName + " " + opts.LastName)
}

type MyFuncOpts1 struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	t1 := MyFuncOpts1{
		"Hemant",
		"Gupta",
		43,
	}
	MyFunc1(t1)
}
