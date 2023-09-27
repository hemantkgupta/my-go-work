package main

import (
	"fmt"
	"strconv"
)

func add2(i int, j int) int { return i + j }

func sub2(i int, j int) int { return i - j }

func mul2(i int, j int) int { return i * j }

func div2(i int, j int) int { return i / j }

type opFuncType2 func(int, int) int

// Functions are values
var opMap2 = map[string]opFuncType2{
	"+": add2,
	"-": sub2,
	"*": mul2,
	"/": div2,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap2[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
