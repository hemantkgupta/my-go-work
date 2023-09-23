package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}

	// Make y slice
	y := make([]int, 4)
	fmt.Println(y, len(y), cap(y))

	// Returns the number of elements copied
	num := copy(y, x)
	fmt.Println(y, num)

	// Use slice expression to copy a range
	z := make([]int, 2)
	num2 := copy(z, x[2:])
	fmt.Println(z, num2)

	// Within slice copy
	num3 := copy(x[:3], x[1:])
	fmt.Println(x, num3)

}
