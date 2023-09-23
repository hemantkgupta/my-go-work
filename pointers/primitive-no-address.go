package main

import "fmt"

func main() {

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName: "Pat",
		// This line won't compile
		// MiddleName: "Perry",
		MiddleName: makePointer("Perry"),
		LastName:   "Peterson",
	}

	fmt.Println(p.FirstName)
}

func makePointer[T any](t T) *T {
	return &t
}
