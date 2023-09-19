package main

import "fmt"

func main() {
	s := "éक्षिaπ囧"
	for i, rn := range s {
		fmt.Printf("%2v: 0x%x %v \n", i, rn, string(rn))
	}
	fmt.Println(len(s))
	right()
}

func right() {
	s := "éक्षिaπ囧"
	for i := 0; i < len(s); i++ {
		fmt.Printf("The byte at index %v: 0x%x \n", i, s[i])
	}
}
