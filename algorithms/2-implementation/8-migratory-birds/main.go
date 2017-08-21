package main

import "fmt"

func main() {
	var (
		n, max int
		birds  []int
		count  [5]int
	)

	fmt.Scan(&n)

	birds, _ = intScanln(n)

	for _, birdType := range birds {
		count[birdType-1]++
	}

	for i := range count {
		if count[i] > count[max] {
			max = i
		}
	}

	fmt.Print(max + 1)
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
