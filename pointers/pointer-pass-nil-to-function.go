package main

import "fmt"

func main() {
	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) // prints nil
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}
