package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	i, c, sum := [...]int{1, 1}, 0, 0
	for c < 4e6 {
		c = i[0] + i[1]
		i[0] = i[1]
		i[1] = c
		if c%2 == 0 {
			sum += c
		}
	}
	fmt.Println("Execution time: ", time.Since(start))
	fmt.Print(sum, "\n")
}
