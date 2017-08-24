package main

import "fmt"

func main() {
	var (
		t         int
		testCases []uint64
	)

	fmt.Scan(&t)

	testCases = make([]uint64, t, t)

	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, n := range testCases {
		fmt.Println(largestPrimeFactor(n))
	}
}

func largestPrimeFactor(n uint64) uint64 {
	var (
		i, max uint64
	)

	i = 2
	max = 2

	for i*i <= n {
		for n%i == 0 {
			max = i
			n /= i
		}

		i++
	}

	if n > max {
		max = n
	}

	return max
}
