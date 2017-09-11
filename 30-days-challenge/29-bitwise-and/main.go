package main

import "fmt"

func main() {
	var (
		t, n, k, max int
	)

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		fmt.Scan(&k)

		max = 0

		for a := 1; a <= n-1; a++ {
			for b := a + 1; b <= n; b++ {
				aAndB := a & b
				if aAndB > max && aAndB < k {
					max = aAndB
				}
			}
		}

		fmt.Println(max)
	}
}
