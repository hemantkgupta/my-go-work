package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	// func json.NewEncoder(w io.Writer) *json.Encoder
	// NewEncoder returns a new encoder that writes to w.
	//
	// func (*json.Encoder).Encode(v any) error
	// Encode writes the JSON encoding of v to the stream,
	// followed by a newline character.
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	var fromFile Person
	// func json.NewDecoder(r io.Reader) *json.Decoder
	// NewDecoder returns a new decoder that reads from r.
	//
	// func (*json.Decoder).Decode(v any) error
	// Decode reads the next JSON-encoded value from its input
	// and stores it in the value pointed to by v.

	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", fromFile)

}
