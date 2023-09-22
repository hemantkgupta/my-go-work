package main

import (
	"fmt"
	"go/format"
)

func main() {
	source := []byte("package p\n\nconst _ = 0000000123i\n")
	output, err := format.Source(source)
	if err != nil {
		fmt.Println("error formatting:", err)
	}
	fmt.Println(string(output))

}
