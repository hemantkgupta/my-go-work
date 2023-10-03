package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "hemantkgupta"

	hash := sha256.New()
	hash.Write([]byte(s))
	binary := hash.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", binary)
}
