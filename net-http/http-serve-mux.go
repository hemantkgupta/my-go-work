package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// func http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.HandleFunc("/", HelloHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}
