package main

import (
	"fmt"
)

func main() {

	//  Array to slice conversion use [:]
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	fmt.Println(xArray)
	fmt.Println(xSlice)

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
