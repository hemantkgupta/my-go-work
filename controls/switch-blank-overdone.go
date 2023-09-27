package main

import (
	"fmt"
	"math/rand"
)

func main() {
	overdoneBlankSwitch()
}

func overdoneBlankSwitch() {
	fmt.Println("This blank switch would be better implemented as a regular switch")

	for i := 0; i < 10; i++ {
		a := rand.Intn(10)

		switch {
		case a == 2:
			fmt.Println("a is 2")
		case a == 3:
			fmt.Println("a is 3")
		case a == 4:
			fmt.Println("a is 4")
		default:
			fmt.Println("a is ", a)
		}
	}

	fmt.Println("\njust use a regular switch like this")
	for i := 0; i < 10; i++ {
		a := rand.Intn(10)

		switch a {
		case 2:
			fmt.Println("a is 2")
		case 3:
			fmt.Println("a is 3")
		case 4:
			fmt.Println("a is 4")
		default:
			fmt.Println("a is ", a)
		}
	}

}
