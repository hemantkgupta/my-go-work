package main

import "fmt"

func main() {
	basicSwitch1()
	fmt.Println()
	basicSwitch2()
	fmt.Println()
	missingLabel()
	fmt.Println()
	labeledBreak()
}

func basicSwitch1() {
	fmt.Println("basic switch statement example")

	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}

	for _, word := range words {

		// Switch based on equality of size
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func basicSwitch2() {
	fmt.Println("basic switch statement example")

	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}

	for _, word := range words {
		size := len(word)
		// Switch based on equality of size
		switch size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}
