package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Slice as buffer
	data := make([]byte, 100)

	for {
		count, err := file.Read(data)
		process(data[:count])
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}
		if count == 0 {
			fmt.Println("All read")
		}
	}
}

func process(buffer []byte) {
	fmt.Println(buffer)
}
