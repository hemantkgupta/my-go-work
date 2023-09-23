package main

import "fmt"

func main() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	fmt.Println(xSlice)
	fmt.Println(xArray)

	// Create small size array
	smallArray := [2]int(xSlice)
	fmt.Println(smallArray)

	// The storage is not shared
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
}
