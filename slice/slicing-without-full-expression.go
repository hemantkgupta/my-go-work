package main

import "fmt"

func main() {

	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")

	// Slicing without using full expression
	// Unused capacity in the original slice is also shared with sub-slices
	y := x[:2]
	z := x[2:]

	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
