package common

// IsPrime returns true if n is a prime number
func IsPrime(n int) bool {
	// 1 is not considered a prime number
	if n == 1 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}
