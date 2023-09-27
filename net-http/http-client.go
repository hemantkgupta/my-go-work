package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "http://webcode.me", nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in response")
	}

	//body, err := io.ReadAll(res.Body)
	// func io.Copy(dst io.Writer, src io.Reader) (written int64, err error)
	// Copy copies from src to dst until either EOF is reached on src or an error occurs.
	// It returns the number of bytes copied and the first error encountered while copying, if any.
	_, err = io.Copy(os.Stdout, res.Body)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))

}
