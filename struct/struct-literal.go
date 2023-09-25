package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
		pet  string
	}

	// Empty struct literal
	bob := person{}

	fmt.Println("The name is: ", bob.name)

	// Struct literal with body
	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println("The name is: ", julia.name)

	// Another style of struct literal with body
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println("The name is: ", beth.name)

}
