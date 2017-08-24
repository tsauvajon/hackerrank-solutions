package main

import (
	"fmt"
	"strconv"
)

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
		fmt.Println(largestPalindrome(n))
	}
}

func largestPalindrome(n int) int {
	max := 101101

	for i := 0; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			if i*j >= n {
				break
			}

			if i*j > max && isPalindrome(i*j) {
				max = i * j
			}
		}
	}

	return max
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
