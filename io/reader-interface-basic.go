package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println("Error in reading letters in string: ", err)
	}
	fmt.Println(counts)
}

func countLetters(r io.Reader) (map[string]int, error) {
	// Considering we are reading one byte for one char at once
	buf := make([]byte, 1)
	out := map[string]int{}
	for {
		// Read data into 2048 byte buffer
		n, err := r.Read(buf)

		// For each byte int buffer
		for _, b := range buf[:n] {
			// Check the content of the byte
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				// Add 1 into the mapped value of string(byte)
				out[string(b)]++
			}
		}
		// The byte read give EOF
		if err == io.EOF {
			return out, nil
		}
		// For any other error
		if err != nil {
			return nil, err
		}
	}
}
