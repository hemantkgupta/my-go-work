package main

import (
	"fmt"
	"slices"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}

	//s := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false

	//fmt.Println(slices.Equal(x, s)) // does not compile
}
