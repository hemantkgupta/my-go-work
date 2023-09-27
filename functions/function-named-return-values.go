package main

import (
	"errors"
	"fmt"
)

func main() {
	x, y, z := divAndRemainder2(5, 2)
	fmt.Println(x, y, z)
}

func divAndRemainder2(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}
