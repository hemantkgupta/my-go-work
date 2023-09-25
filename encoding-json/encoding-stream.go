package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)

	type per struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	t1 := per{}
	t1.Name = "Hemant"
	t1.Age = 43

	t2 := per{
		"Ram",
		5,
	}

	allInputs := [2]per{t1, t2}

	for _, input := range allInputs {
		err := enc.Encode(input)
		if err != nil {
			panic(err)
		}
		fmt.Println(b.String())
	}
	//out := b.String()
}
