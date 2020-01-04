package challenge

// LargestPrimeFactor returns the largest prime factor of the given number
func LargestPrimeFactor(number int) int {
	primeFactors := []int{}
	// Take the lowest prime number to begin
	prime := 2

	for prime <= number {
		// If the number is divisible by the prime, keep going
		canDivide := number%prime == 0
		if canDivide {
			number = number / prime
			primeFactors = append(primeFactors, prime)
		} else {
			prime = getNextPrime(prime)
		}
	}
	return primeFactors[len(primeFactors)-1]
}

func getNextPrime(n int) int {
	i := n + 1
	for {
		if IsPrime(i) {
			return i
		}
		i++
	}
}

// IsPrime returns true if n is a prime number
func IsPrime(n int) bool {
	// 1 is not considered a prime number
	if n == 1 {
		return false
	}

	for i := 2; i < n/2; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}
