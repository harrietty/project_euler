package challenge

// MultiplesOf3and5 finds the sum of all the multiples of 3 or 5 below the provided parameter value number.
func MultiplesOf3and5(number int) int {
	total := 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}
