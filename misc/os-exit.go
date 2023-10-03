package main

import (
	"fmt"
	"os"
)

func main() {
	// Defer will not be executed if os.Exit is called.
	defer fmt.Println("!")

	os.Exit(3)
}
