package main

import "fmt"

func main() {
	var (
		t         int
		testCases []int
	)

	fmt.Scan(&t)

	testCases = make([]int, t)

	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, n := range testCases {
		fmt.Println(smallestMultiple(n))
	}
}

func smallestMultiple(n int) int {
	if n == 1 {
		return 1
	}

	current := 2

	for !isDivisibleByRange(n, current) {
		current += 2
	}

	return current
}

func isDivisibleByRange(n, current int) bool {
	for i := 2; i <= n; i++ {
		if current%i != 0 {
			return false
		}
	}

	return true
}
