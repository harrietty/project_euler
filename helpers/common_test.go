package common

import "testing"

func TestIsPrime(t *testing.T) {
	result := IsPrime(13195)
	expected := false
	if result != expected {
		t.Errorf("TestIsPrime was incorrect, got: %v, want: %v.", result, expected)
	}
}
