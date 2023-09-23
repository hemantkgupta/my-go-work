package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
		pet  string
	}

	bob := person{}

	julia := person{
		"Julia",
		40,
		"cat",
	}

	fmt.Println(bob.name)
	fmt.Println(julia.name)
}
