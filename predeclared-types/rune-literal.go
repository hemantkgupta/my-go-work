package main

import "fmt"

func main() {

	// All these rune literals represent 97 in decimal
	a := 'a'    // single Unicode characters
	b := '\141' // 8-bit octal numbers
	c := '\x61' // 8-bit hexadecimal numbers

	d := '\u0061' // 16-bit hexadecimal numbers

	e := '\U00000061' // 32-bit Unicode numbers ()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
