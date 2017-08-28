package main

import "fmt"

func main() {
	var (
		t         int
		testCases []uint64
	)

	fmt.Scan(&t)

	testCases = make([]uint64, t)

	for i := 0; i < t; i++ {
		fmt.Scan(&testCases[i])
	}

	for _, n := range testCases {
		fmt.Println(evenFibonacciSum(n))
	}
}

func evenFibonacciSum(n uint64) uint64 {
	var (
		sum, a, b uint64
	)

	sum = 0
	a = 1
	b = 2

	for b < n {
		if b%2 == 0 {
			sum += b
		}

		a, b = b, a+b
	}

	return sum
}
