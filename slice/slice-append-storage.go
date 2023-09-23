package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}

	// y is same storage as x
	y := x[:2]
	fmt.Println(cap(x), cap(y))

	// Append to y updates the storage
	// Now x also have updated value
	y = append(y, "z")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
