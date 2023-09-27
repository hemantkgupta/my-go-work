package main

import "fmt"

func main() {

	// There is no name for the function
	// Anonymous function
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}

	for i := 0; i < 5; i++ {
		f(i)
	}

	for i := 0; i < 5; i++ {
		// You don't need to assign a vairable for the function
		// Create and invoke immediately
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}
