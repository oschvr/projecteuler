package main

import (
	"fmt"
	"math"
	"time"
)

const n int64 = 600851475143

func isPrime(x int64) bool {
	var i int64 = 2
	for ; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	var i = int64(math.Sqrt(float64(n)))
	for ; i > 1; i-- {
		if n%i == 0 && isPrime(i) {
			fmt.Println("Program took: ", time.Since(start))
			fmt.Println("Largest Prime Factor: ", i)
			break
		}
	}
}
