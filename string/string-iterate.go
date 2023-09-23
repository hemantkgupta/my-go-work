package main

import "fmt"

func main() {
	sample := "apple_π!"
	// When for loop sees a multibyte rune in a string,
	// it converts the UTF-8 representation into a single 32-bit number
	// and assigns it to the value.
	// The offset is incremented by the number of bytes in the rune

	// It iterates over the runes, not the bytes.
	for i, r := range sample {
		fmt.Println(i, r, string(r))
	}
	fmt.Println()

}
