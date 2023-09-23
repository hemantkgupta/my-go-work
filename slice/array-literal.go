package main

import "fmt"

func main() {
	// Declaration
	var x [3]int
	fmt.Println(x) // [0 0 0]

	//
	var y = [3]int{10, 20, 30}
	fmt.Println(y) // [10 20 30]

	// Sparse array
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z) // [1 0 0 0 0 4 6 0 0 0 100 15]

	// Array literal ... notation for size
	var a = [...]int{10, 20, 30}
	fmt.Println(a) //
}
