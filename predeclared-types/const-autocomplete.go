package main

import "fmt"

func main() {
	const (
		X float32 = 3.14
		Y         // here must be one identifier
		Z         // here must be one identifier

		A, B = "Go", "language"
		C, _
		// In the above line, the blank identifier
		// is required to be present.
	)

	fmt.Println(X, Y, Z, A, B, C)
}
