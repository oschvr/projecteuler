package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	largestPalin := 0
	a := 0
	for ; a <= 999; a++ {
		b := 100
		for ; b <= 999; b++ {
			c := a * b
			if isPalindrome(c) && c > largestPalin {
				largestPalin = c
			}
		}
	}
	fmt.Println("Execution time: ", time.Since(start))
	fmt.Println("Largest Palindrome: ", largestPalin)
}

func reversed(n int) int { // 456789
	reversed := 0 // 0
	for ; n > 0; n = n / 10 {
		reversed = 10*reversed + n%10
	}
	return reversed
}

func isPalindrome(n int) bool {
	return n == reversed(n)
}
