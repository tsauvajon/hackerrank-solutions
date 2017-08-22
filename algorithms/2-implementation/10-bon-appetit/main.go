package main

import "fmt"

func main() {
	var (
		n, k                          int
		total, expected, actual, diff float64
		params, prices                []int
	)

	params, _ = intScanln(2)

	n = params[0]
	k = params[1]

	prices, _ = intScanln(n)

	fmt.Scan(&actual)

	for i, val := range prices {
		if i != k {
			total += float64(val)
		}
	}

	expected = total / 2

	diff = actual - expected

	if diff > 0 {
		fmt.Printf("%.0f", diff)
		return
	}

	fmt.Print("Bon Appetit")
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
