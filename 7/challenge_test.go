package challenge

import "testing"

func TestNthPrime(t *testing.T) {
	actual := NthPrime(10)
	expected := 29
	if actual != expected {
		t.Errorf("TestNthPrime was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = NthPrime(1000)
	expected = 7919
	if actual != expected {
		t.Errorf("TestNthPrime was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = NthPrime(10001)
	expected = 104743
	if actual != expected {
		t.Errorf("TestNthPrime was incorrect, got: %d, want: %d.", actual, expected)
	}
}
