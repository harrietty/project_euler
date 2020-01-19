package challenge

import "testing"

func TestLargestProductInSeries(t *testing.T) {
	actual := LargestProductInSeries(4)
	expected := 5832
	if actual != expected {
		t.Errorf("TestLargestProductInSeries was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = LargestProductInSeries(13)
	expected = 23514624000
	if actual != expected {
		t.Errorf("TestLargestProductInSeries was incorrect, got: %d, want: %d.", actual, expected)
	}
}
