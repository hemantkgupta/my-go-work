package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	runes := "🐜🐬🐄🐘🦂🐫🐑🦍🐯🐞"

	f, err := os.Create("runes.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	wr := bufio.NewWriter(f)

	for _, _rune := range runes {

		wr.WriteRune(_rune)
		wr.WriteRune('\n')
	}

	wr.Flush()

	fmt.Println("runes written")
}
