package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < len(m); i++ {
		fmt.Println("Loop", i)

		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
