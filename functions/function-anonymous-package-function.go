package main

import "fmt"

var (
	add1 = func(i, j int) int { return i + j }
)

func main() {
	x := add1(2, 3)
	fmt.Println(x)

	changeAdd()
	y := add1(2, 3)
	fmt.Println(y)
}

func changeAdd() {
	// Assign a new value to
	// a package-level anonymous function:
	add1 = func(i, j int) int {
		return i + j + j
	}
}
