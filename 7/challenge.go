package challenge

import common "github.com/harrietty/project_euler/helpers"

// NthPrime finds the nth prime number
func NthPrime(n int) int {
	p := 0

	// 2 is the first prime number
	i := 2
	for {
		if common.IsPrime(i) {
			p++
			if p == n {
				return i
			}
		}
		i++
	}
}
