package main

import "fmt"

func main() {
	// String to byte slice
	hello := []byte("Hello ")
	world := "world!"

	// The normal way:
	// helloWorld := append(hello, []byte(world)...)
	helloWorld := append(hello, world...) // sugar way
	fmt.Println(string(helloWorld))

	helloWorld2 := make([]byte, len(hello)+len(world))
	copy(helloWorld2, hello)
	fmt.Println(string(helloWorld2))

	// The normal way:
	// copy(helloWorld2[len(hello):], []byte(world))
	copy(helloWorld2[len(hello):], world) // sugar way
	fmt.Println(string(helloWorld2))
}
