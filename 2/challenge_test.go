package challenge

import (
	"testing"
)

func TestEvenFibSum(t *testing.T) {
	result := EvenFibSum(10)
	expected := 44
	if result != expected {
		t.Errorf("TestEvenFibSum was incorrect, got: %d, want: %d.", result, expected)
	}

	result = EvenFibSum(43)
	expected = 350704366
	if result != expected {
		t.Errorf("TestEvenFibSum was incorrect, got: %d, want: %d.", result, expected)
	}
}
