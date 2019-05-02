package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var a, i int64 = 1, 2
	for ; i <= 20; i++ {
		a = lcm(a, i)
	}
	fmt.Println("Execution time: ", time.Since(start))
	fmt.Println("Smallest Multiple: ", a)
}

func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

func gcd(a, b int64) int64 {
	c := a % b
	if c == 0 {
		return b
	}
	return gcd(b, c)
}
