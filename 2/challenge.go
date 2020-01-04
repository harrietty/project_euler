package challenge

// EvenFibSum finds the sum of the even-valued terms by considering the terms in the Fibonacci sequence that do not exceed the nth term
func EvenFibSum(number int) int {
	total := 2
	n1 := 1
	n2 := 2
	fib := 0
	count := 2
	for count < number {
		count++
		fib = n1 + n2
		n1 = n2
		n2 = fib
		if fib%2 == 0 {
			total += fib
		}
	}

	return total
}
