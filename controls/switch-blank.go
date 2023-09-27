package main

import (
	"fmt"
)

func main() {
	basicBlankSwitch1()
	fmt.Println()
	basicBlankSwitch2()
}

func basicBlankSwitch1() {
	fmt.Println("a basic blank switch")

	words := []string{"hi", "salutations", "hello"}

	for _, word := range words {
		wordLen := len(word)

		switch {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}

func basicBlankSwitch2() {
	fmt.Println("a basic blank switch")

	words := []string{"hi", "salutations", "hello"}

	for _, word := range words {

		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}
