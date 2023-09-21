package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	testValues := []int{10, 20, 30}
	_, sum := sortAndTotal(testValues)
	expected := 60
	if sum != expected {
		t.Fatalf("Expected %v, Got %v", expected, sum)
	}
}

func TestSort(t *testing.T) {
	testValues := []int{1, 279, 48, 12, 3}
	// Use in-built IntsAreSorted method for testing
	sorted, _ := sortAndTotal(testValues)
	if !sort.IntsAreSorted(sorted) {
		t.Fatalf("Unsorted data %v", sorted)
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 250
	data := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Fill the random data
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		b.StartTimer()
		sortAndTotal(data)
	}
}
