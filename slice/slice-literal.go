package main

import "fmt"

func main() {
	// Empty slice literal - It is not nil slice
	var a = []int{}
	fmt.Println(a)

	// Slice literal with 3 values
	var x = []int{10, 20, 30}
	fmt.Println(x) // [10 20 30]

	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y) // [1 0 0 0 0 4 6 0 0 0 100 15]

}
