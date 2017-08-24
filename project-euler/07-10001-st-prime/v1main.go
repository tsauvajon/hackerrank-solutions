package main

import "fmt"

func main() {
	var (
		t         int
		testCases []int
	)

	fmt.Scan(&t)

	testCases = make([]int, t, t)

	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, n := range testCases {
		fmt.Println(nthPrime(n))
	}
}

func nthPrime(n int) uint {
	var m uint

	for i := 0; i < n; i++ {
		m = nextPrime(m)
	}

	return m
}

func nextPrime(n uint) uint {
	n++

	for largestPrimeFactor(n) != n {
		n++
	}

	return n
}

func largestPrimeFactor(n uint) uint {
	var (
		i, max uint
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
