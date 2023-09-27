package main

import (
	"io"
	"log"
	"os"
)

func main() {
	/*
		dir, err := filepath.Abs("./")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)

		f, err := os.Open("./test.txt")

		if len(os.Args) < 2 {
			log.Fatal("no file specified")
		}
	*/
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
