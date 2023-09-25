package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		fmt.Println("Error in reading the file: ", err)
	}
	defer closer()
	counts, err := countLetters2(r)
	if err != nil {
		fmt.Println("Error in counting letters: ", err)
	}
	fmt.Println(counts)
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func countLetters2(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
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
