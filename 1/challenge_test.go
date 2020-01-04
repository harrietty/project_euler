package challenge

import "testing"

func TestMultiplesOf3and5(t *testing.T) {
	total := MultiplesOf3and5(49)
	expected := 543
	if total != expected {
		t.Errorf("MultiplesOf3and5 was incorrect, got: %d, want: %d.", total, expected)
	}

	total = MultiplesOf3and5(1000)
	expected = 233168
	if total != expected {
		t.Errorf("MultiplesOf3and5 was incorrect, got: %d, want: %d.", total, expected)
	}

	total = MultiplesOf3and5(19564)
	expected = 89301183
	if total != expected {
		t.Errorf("MultiplesOf3and5 was incorrect, got: %d, want: %d.", total, expected)
	}

	total = MultiplesOf3and5(8456)
	expected = 16687353
	if total != expected {
		t.Errorf("MultiplesOf3and5 was incorrect, got: %d, want: %d.", total, expected)
	}
}
