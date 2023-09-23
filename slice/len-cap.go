package main

import "fmt"

func main() {
	// Initially both len and cap are zero
	var x []int
	fmt.Println(x, len(x), cap(x))

	// append will increae the len
	// Also if needed cap will be increased
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}
