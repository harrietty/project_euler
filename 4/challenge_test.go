package challenge

import "testing"

func TestLargestPalindromeProduct(t *testing.T) {
	result := LargestPalindromeProduct(2)
	expected := 9009
	if result != expected {
		t.Errorf("TestLargestPalindromeProduct was incorrect, got: %d, want: %d.", result, expected)
	}

	result = LargestPalindromeProduct(3)
	expected = 906609
	if result != expected {
		t.Errorf("TestLargestPalindromeProduct was incorrect, got: %d, want: %d.", result, expected)
	}
}
