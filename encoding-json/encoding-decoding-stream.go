package main

import (
	"encoding/json"
	"strings"
)

func main() {
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		//
	}
}
