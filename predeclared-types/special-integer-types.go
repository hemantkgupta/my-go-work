package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// byte is an alias for uint8
	var a byte = 123
	fmt.Printf("type of a %T\n", a)

	// int - machine dependent
	var b int = 123
	fmt.Printf("type of a %T\n", b)

	// uint
	var c int = 123
	fmt.Printf("type of a %T\n", c)

	// rune is an alias for int32
	var d rune = 123
	fmt.Printf("type of a %T\n", d)

	//uintptr
	s := &sample{
		a: 1,
		b: "test",
	}

	//Getting the address of field b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	//Typecasting it to a string pointer and printing the value of it
	fmt.Println(*(*string)(p))
}

type sample struct {
	a int
	b string
}
