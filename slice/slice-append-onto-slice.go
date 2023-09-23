package main

import "fmt"

func main() {
	var x = []int{1, 2, 3}
	// Single item appended
	x = append(x, 4)
	fmt.Println(x)

	// Multiple items appended
	x = append(x, 5, 6, 7)
	fmt.Println(x)

	y := []int{20, 30, 40}
	// Another slice is appended
	x = append(x, y...)
	fmt.Println(x)

}
