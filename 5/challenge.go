package challenge

// SmallestMult returns the smallest positive number that is evenly divisible by all of the numbers from 1 to n
func SmallestMult(n int) int {
	// i: starting at n, keep trying numbers
	i := n
	for {
		isDivisibleByAll := true
		for j := 1; j <= n; j++ {
			if i%j != 0 {
				isDivisibleByAll = false
				break
			}
		}
		if isDivisibleByAll {
			return i
		}
		i++
	}
}
