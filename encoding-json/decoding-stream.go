package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// Notice there is no comma after first json struct
	// It is representing the stream
	data := `{
        "name": "John Doe",
        "age": 24
    }
	{
        "name": "Brian Flemming",
        "age": 42
    }`

	dec := json.NewDecoder(strings.NewReader(data))

	// func (*json.Decoder).More() bool
	// More reports whether there is another element in the current array or object being parsed.
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t.Name)
	}
}
