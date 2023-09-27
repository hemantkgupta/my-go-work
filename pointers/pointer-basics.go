package main

import "fmt"

func main() {
	// A pointer is a variable.
	// Its content is the address where another variable is stored.

	// The & is the address operator. It precedes a value type.
	// It returns the address where the value is stored.

	// A pointer type is a type that represents a pointer.
	// It is written with a * before a type name.

	// The * is the indirection operator. It precedes a variable of pointer type.
	// It returns the pointed-to value.
	// This is called de-referencing.

	var x = 10
	pointerToX := &x
	fmt.Println("The address of x is: ", &x)
	fmt.Println("The value of x is: ", *pointerToX)

	y := 20
	var pointerToY *int
	pointerToY = &y
	fmt.Println("The address of y is: ", &y)
	fmt.Println("The value of y is: ", *pointerToY)

	z := 5 + *pointerToX
	fmt.Println("The value of z is: ", z)

	// Program will panic if you attempt to dereference a nil pointer
	var a *int
	fmt.Println(a == nil)
	//fmt.Println(*y)

	// The built-in function new creates a pointer variable.
	// It returns a pointer to a zero value instance of the provided type
	// The new function is rarely used.
	var b = new(int)
	fmt.Println(b == nil) // prints false
	fmt.Println(*b)       // prints 0

}
