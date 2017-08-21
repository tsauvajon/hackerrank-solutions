package main

import "fmt"

func main() {
	var (
		n, m, d, count int
	)

	fmt.Scan(&n)

	values, _ := intScanln(n)
	date, _ := intScanln(2)

	d = date[0]
	m = date[1]

	for i := 0; i <= n-m; i++ {
		sum := 0

		for j := 0; j < m; j++ {
			sum += values[j+i]
		}

		if sum == d {
			count++
		}
	}

	fmt.Print(count)
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
