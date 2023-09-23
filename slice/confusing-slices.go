package main

import "fmt"

func main() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")

	// y has first two value
	// It will have same capacity as x
	y := x[:2]
	fmt.Println("y:", y)

	// z has values from index 2 onwards
	// It will have capacity 5-2 = 3
	z := x[2:]
	fmt.Println("z:", z)

	fmt.Println(cap(x), cap(y), cap(z))
	// Appended 3 values to y
	y = append(y, "i", "j", "k")
	fmt.Println("y:", y)

	// Appended 1 value to source x
	// Will this add a new value to x? No!
	x = append(x, "x")
	fmt.Println("x:", x)

	// Now appended 1 value to z
	z = append(z, "y")
	fmt.Println("z:", z)

	fmt.Println("x:", x)
	fmt.Println("y:", y)

}
