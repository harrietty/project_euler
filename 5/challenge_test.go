package challenge

import "testing"

func TestSmallestMult(t *testing.T) {
	actual := SmallestMult(5)
	expected := 60
	if actual != expected {
		t.Errorf("TestSmallestMult was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = SmallestMult(7)
	expected = 420
	if actual != expected {
		t.Errorf("TestSmallestMult was incorrect, got: %d, want: %d.", actual, expected)
	}

	actual = SmallestMult(13)
	expected = 360360
	if actual != expected {
		t.Errorf("TestSmallestMult was incorrect, got: %d, want: %d.", actual, expected)
	}
}
