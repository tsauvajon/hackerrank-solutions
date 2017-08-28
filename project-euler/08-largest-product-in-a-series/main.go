package main

import (
	"fmt"
)

func main() {
	var (
		t int
	)

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var (
			n, k    int
			integer string
		)

		fmt.Scan(&n)
		fmt.Scan(&k)
		fmt.Scanln(&integer)

		fmt.Println(greatestProduct(k, integer))
	}
}

func greatestProduct(k int, integer string) int {
	var (
		max, current int
	)

	for i := 0; i < len(integer)-k; i++ {
		current = 1
		for j := 0; j < k; j++ {
			n := integer[j+i] - '0'
			current *= int(n)
		}

		if current > max {
			max = current
		}
	}

	return max
}
