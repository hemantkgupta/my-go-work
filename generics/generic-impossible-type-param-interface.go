package main

import "fmt"

type ImpossiblePrintableInt interface {
	int
	String() string
}

type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	val T
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprint(mi)
}

func main() {
	// int does not satisfy ImpossiblePrintableInt (missing method String)
	//s := ImpossibleStruct[int]{10}

	// MyInt does not satisfy ImpossiblePrintableInt (possibly missing ~ for int in ImpossiblePrintableInt)
	// s2 := ImpossibleStruct[MyInt]{10}
}
