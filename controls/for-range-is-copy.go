package main

import "fmt"

func main() {
	forRangeIsACopy()
}

func forRangeIsACopy() {
	fmt.Println("for-range loop, show that modifying value variable doesn't modify the original slice")
	evenVals := []int{2, 4, 6, 8, 10, 12}

	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}
