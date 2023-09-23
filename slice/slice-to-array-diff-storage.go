package main

import "fmt"

func main() {
	xSlice := []int{1, 2, 3, 4}

	// All elements of slice are copied
	// to a new storage
	xArray := [4]int(xSlice)
	fmt.Println(xArray)

	// Only given number of elements are copied
	// The size of the array can be smaller than the size of the slice
	smallArray := [2]int(xSlice)
	fmt.Println(smallArray)

	// Update the slice
	xSlice[0] = 10
	fmt.Println(xSlice)

	// No change in arrays created
	fmt.Println(xArray)
	fmt.Println(smallArray)
}
