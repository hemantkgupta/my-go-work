package main

import "fmt"

func main() {
	var s string = "Hello, 🌞"
	n := len(s)
	fmt.Println("No of bytes in string:", n)
	var bs []byte = []byte(s)
	fmt.Println(bs)

	// Rune size will not be equal to n
	var rs []rune = []rune(s)
	fmt.Println("No of runes in string:", len(rs))
	fmt.Println(rs)
}
