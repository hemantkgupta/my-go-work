package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// func http.Get(url string) (resp *http.Response, err error)
	//
	// Get issues a GET to the specified URL.
	// If the response is one of the following redirect codes, Get follows the redirect,
	//  up to a maximum of 10 redirects:
	// 301 (Moved Permanently)
	// 302 (Found)
	// 303 (See Other)
	// 307 (Temporary Redirect)
	// 308 (Permanent Redirect)
	// An error is returned if there were too many redirects or if there was an HTTP protocol error.
	resp, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)

	fmt.Println(string(body))
}
