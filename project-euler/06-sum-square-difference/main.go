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
		fmt.Println(sumSquareDifference(n))
	}
}

func sumSquareDifference(n int) int {
	var (
		squareSum, sumSquares int
	)

	for i := 1; i <= n; i++ {
		squareSum += i
		sumSquares += i * i
	}

	squareSum *= squareSum

	return squareSum - sumSquares
}
