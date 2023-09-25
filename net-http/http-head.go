package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// func http.Head(url string) (resp *http.Response, err error)
	// Head issues a HEAD to the specified URL.
	// If the response is one of the redirect codes,
	// Head follows the redirect, up to a maximum of 10 redirects:
	// Head is a wrapper around DefaultClient.Head.

	resp, err := http.Head("http://webcode.me")

	if err != nil {

		log.Fatal(err)
	}

	for k, v := range resp.Header {

		fmt.Printf("%s %s\n", k, v)
	}
}
