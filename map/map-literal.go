package main

import "fmt"

func main() {
	// Empty map literal - Remember the syntax
	totalWins := map[string]int{}

	fmt.Println(totalWins)

	// Map literal with body
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
}
