package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n, t int
	)

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)

		if isNPrime(n) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func isNPrime(n int) bool {
	if n == 1 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
