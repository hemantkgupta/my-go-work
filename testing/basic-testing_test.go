package main

import "testing"

func TestAddNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Errorf("incorrect result: expected %d, got %d", 5, result)
	}
}
