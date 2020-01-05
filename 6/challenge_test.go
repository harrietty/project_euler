package challenge

import "testing"

func TestSumSquareDifference(t *testing.T) {
	actual := SumSquareDifference(10)
	expected := 2640
	if actual != expected {
		t.Errorf("TestSumSquareDifference was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = SumSquareDifference(20)
	expected = 41230
	if actual != expected {
		t.Errorf("TestSumSquareDifference was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = SumSquareDifference(100)
	expected = 25164150
	if actual != expected {
		t.Errorf("TestSumSquareDifference was incorrect, got: %d, want: %d.", actual, expected)
	}
}
