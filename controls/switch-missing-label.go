package main

import "fmt"

func main() {
	missingLabel()
}

func missingLabel() {
	fmt.Println("The case of the missing label...")

	for i := 0; i < 10; i++ {

		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			// Interpreted as exiting the case
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
}
