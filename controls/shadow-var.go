package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		// This will shadow
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}
