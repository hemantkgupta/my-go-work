package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker4(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker4("not_here.txt")
	if err != nil {
		// Compare with wrapped error
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}
}
