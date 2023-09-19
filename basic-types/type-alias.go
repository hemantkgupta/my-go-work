package main

import "fmt"

func main() {
	type boolean = bool
	type Text = string
	type U8 = uint8
	type char = rune

	var x boolean
	fmt.Println(x)

	var y Text
	y = "Test String"
	fmt.Println(y)

	var z U8
	z = 11
	fmt.Println(z)

	var a char
	a = 'h'
	fmt.Println(a)
}
