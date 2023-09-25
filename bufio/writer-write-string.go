package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := []string{"an old falcon", "misty mountains",
		"a wise man", "a rainy morning"}

	wr := bufio.NewWriter(os.Stdout)

	for _, line := range data {
		// Write to buffer
		wr.WriteString(line + "\n")
	}
	// Write to target from buffer
	wr.Flush()

	fmt.Println("data written")
}
