package main

import "testing"

var result int

func BenchmarkCalculate(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = Calculate(5)
	}
	result = r
}
