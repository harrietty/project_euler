package challenge

import "testing"

func TestPythagoreanTriplet(t *testing.T) {
	actual := PythagoreanTriplet(24)
	expected := 480
	if actual != expected {
		t.Errorf("TestPythagoreanTriplet was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = PythagoreanTriplet(1000)
	expected = 31875000
	if actual != expected {
		t.Errorf("TestPythagoreanTriplet was incorrect, got: %d, want: %d.", actual, expected)
	}
}
