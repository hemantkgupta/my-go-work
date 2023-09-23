package main

import "fmt"

func main() {

	x := 10
	failedUpdate(&x)
	// prints 10
	fmt.Println(x)

	update(&x)
	fmt.Println(x) // prints 20
}

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}
