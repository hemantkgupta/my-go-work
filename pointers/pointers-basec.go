package main

import "fmt"

func main() {

	x := 10
	pointerToX := &x

	// Can be written as
	// var pointerToX *int
	// pointerToX = &x

	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(*pointerToX) // prints 10

	z := 5 + *pointerToX
	fmt.Println(z) // prints 15

	var y *int
	fmt.Println(y == nil)
	// panics
	//fmt.Println(*y)

	// The new function is rarely used
	var a = new(int)
	fmt.Println(a == nil) // prints false
	fmt.Println(*a)       // prints 0

}
