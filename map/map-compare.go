package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}

	n := map[string]int{
		"world": 10,
		"hello": 5,
	}

	fmt.Println(maps.Equal(m, n))

}
