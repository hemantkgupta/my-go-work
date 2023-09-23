package main

import (
	"fmt"
)

func main() {
	x := [4]int{5, 6, 7, 8}

	y := x[:2]
	z := x[2:]

	// Update array
	x[0] = 10
	fmt.Println("x:", x)

	// See the slice changed
	fmt.Println("y:", y)

	fmt.Println("z:", z)
}
