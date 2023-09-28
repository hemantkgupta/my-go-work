package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	r, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	f, err := os.Create("index.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// func (*os.File).ReadFrom(r io.Reader) (n int64, err error)
	// ReadFrom implements io.ReaderFrom.

	_, err = f.ReadFrom(r.Body)

	if err != nil {
		log.Fatal(err)
	}
}
