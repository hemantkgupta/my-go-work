package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// func http.PostForm(url string, data url.Values) (resp *http.Response, err error)
	// PostForm issues a POST to the specified URL, with data's keys and values URL-encoded as the request body.
	// The Content-Type header is set to application/x-www-form-urlencoded.
	// To set other headers, use NewRequest and DefaultClient.Do.
	resp, err := http.PostForm("https://httpbin.org/post",
		url.Values{"name": {"John Doe"}, "message": {"Hey!"}})

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
