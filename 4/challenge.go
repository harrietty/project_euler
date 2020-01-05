package challenge

import (
	"math"
	"strconv"
)

// LargestPalindromeProduct finds the largest palindrome made from the product of two n-digit numbers
func LargestPalindromeProduct(n int) int {
	max := int(math.Pow(10, float64(n)))
	pal := 0
	for i := 0; i < max; i++ {
		for j := 1; j < max; j++ {
			product := i * j
			if isPalindrome(product) && product > pal {
				pal = product
			}
		}
	}

	return pal
}

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	runes := []rune(str)
	mid := len(str) / 2
	mid2 := mid
	if len(str)%2 == 1 {
		mid2 = mid + 1
	}
	substr1 := string(runes[0:mid])
	substr2 := reverse(string(runes[mid2:]))
	return substr1 == substr2
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
