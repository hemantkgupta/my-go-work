package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	words := []string{}

	data := "A foggy mountain.\nAn old falcon.\nA wise man."

	// func bufio.NewScanner(r io.Reader) *bufio.Scanner
	// func strings.NewReader(s string) *strings.Reader
	sc := bufio.NewScanner(strings.NewReader(data))

	// func (*bufio.Scanner).Split(split bufio.SplitFunc)
	// Split sets the split function for the Scanner.
	// The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	//
	// func bufio.ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
	// ScanWords is a split function for a Scanner that returns each space-separated word of text,
	// with surrounding spaces deleted.
	// It will never return an empty string.
	// The definition of space is set by unicode.IsSpace.
	sc.Split(bufio.ScanWords)

	n := 0

	for sc.Scan() {
		words = append(words, sc.Text())
		n++
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# of words: %d\n", n)

	for _, word := range words {

		fmt.Println(word)
	}
}
