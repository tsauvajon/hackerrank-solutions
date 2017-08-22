package main

import "fmt"

func main() {
	var (
		n, min, max, minCount, maxCount int
	)

	fmt.Scan(&n)

	values, _ := intScanln(n)

	min = values[0]
	max = values[0]

	for i := 1; i < n; i++ {
		if values[i] > max {
			max = values[i]
			maxCount++
		}

		if values[i] < min {
			min = values[i]
			minCount++
		}
	}

	fmt.Printf("%d %d", maxCount, minCount)
}

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}
