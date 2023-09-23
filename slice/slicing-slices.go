package main

import "fmt"

func main() {

	x := []string{"a", "b", "c", "d"}

	// Till index 2 (exclusive) onwards
	y := x[:2]

	// From index 1 (inclusive) onwards
	z := x[1:]

	// From 1 (inclusive) to 3 (exclusive)
	d := x[1:3]

	// All
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
