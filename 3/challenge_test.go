package challenge

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	result := LargestPrimeFactor(5)
	expected := 5
	if result != expected {
		t.Errorf("LargestPrimeFactor was incorrect, got: %d, want: %d.", result, expected)
	}

	result = LargestPrimeFactor(13195)
	expected = 29
	if result != expected {
		t.Errorf("LargestPrimeFactor was incorrect, got: %d, want: %d.", result, expected)
	}

	result = LargestPrimeFactor(600851475143)
	expected = 6857
	if result != expected {
		t.Errorf("LargestPrimeFactor was incorrect, got: %d, want: %d.", result, expected)
	}
}
