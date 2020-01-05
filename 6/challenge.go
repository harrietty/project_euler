package challenge

import "math"

// SumSquareDifference finds the difference between the sum of the squares of the first n natural numbers and the square of the sum.
func SumSquareDifference(n int) int {
	sumOfSquares := 0
	sum := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
		sum += i
	}

	return int(math.Abs(float64(sum*sum) - float64(sumOfSquares)))
}
