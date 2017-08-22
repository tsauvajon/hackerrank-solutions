package main

import "fmt"

func main() {
	var (
		n, k, count int
		a           []int
	)

	params, _ := intScanln(2)

	n = params[0]
	k = params[1]

	a, _ = intScanln(n)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (a[i]+a[j])%k == 0 {
				count++
			}
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
