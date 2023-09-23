package main

import "fmt"

func main() {

	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	fmt.Println(xSlice)
	fmt.Println(xArrayPointer)

	// The storage between the two is shared
	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println(xSlice)        // prints [10 20 3 4]
	fmt.Println(xArrayPointer) // prints &[10 20 3 4]

}
