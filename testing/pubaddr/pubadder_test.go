package pubaddr_test

import (
	"github.com/hemantkgupta/my-go-work/gotesting/pubaddr"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := pubaddr.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
