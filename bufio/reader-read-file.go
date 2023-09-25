package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file_example_JPG_500KB.jpg")

	if err != nil {
		log.Fatal("Error in reading file: ", err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)

		if err != nil {

			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		fmt.Printf("%s", hex.Dump(buf))
	}
}
