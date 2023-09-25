package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// A file is passed 
	// func bufio.NewScanner(r io.Reader) *bufio.Scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
