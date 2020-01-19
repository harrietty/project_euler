package challenge

import common "github.com/harrietty/project_euler/helpers"

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
		if common.IsPrime(i) {
			return i
		}
		i++
	}
}
