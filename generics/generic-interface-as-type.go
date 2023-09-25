package main

import "fmt"

func main() {

}

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}
