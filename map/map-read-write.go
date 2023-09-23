package main

import "fmt"

func main() {
	// string to int map
	totalWins := map[string]int{}

	// Put into map
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	// Get from map
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])

	// Update map value
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])

	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	// Delete from map
	delete(totalWins, "Lions")
	fmt.Println(totalWins)

	// Remove all entries from map
	clear(totalWins)
	fmt.Println(totalWins, len(totalWins))

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	// ok idiom
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}
