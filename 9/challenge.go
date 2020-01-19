package challenge

import "math"

// PythagoreanTriplet returns the product of the pythagorean triplet that equals n
func PythagoreanTriplet(n int) int {
	// a2 + b2 = c2   a < b < c
	// start by finding a, start at 1
	for a := 1; a <= n-2; a++ {
		// the value for b must be greater than a, a + 1
		for b := a + 1; b <= n-a-1; b++ {
			// c = n - a - b
			c := n - a - b
			if math.Pow(float64(a), 2.0)+math.Pow(float64(b), 2.0) == math.Pow(float64(c), 2.0) {
				return a * b * c
			}
		}
	}
	return 1
}
