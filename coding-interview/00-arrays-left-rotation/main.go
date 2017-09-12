package main

import "fmt"

func main() {
	var (
		n, d int
		a    []int
	)

	fmt.Scan(&n)
	fmt.Scan(&d)

	a, _ = intScanln(n)

	for i := 0; i < d; i++ {
		a = append(a[1:], a[0])
	}

	for _, val := range a {
		fmt.Printf("%d ", val)
	}
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
