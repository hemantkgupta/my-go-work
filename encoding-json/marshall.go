package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {

	u1 := User{1, "John Doe", "gardener"}

	// func json.Marshal(v any) ([]byte, error)
	// Marshal returns the JSON encoding of v.
	// Marshal traverses the value v recursively.
	json_data_byte_slice, err := json.Marshal(u1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data_byte_slice))

	users := []User{
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
	}

	json_data2, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data2))
}
