package main

import "fmt"

func main() {
	var (
		n, count int
		drawer   map[int]bool
	)

	fmt.Scan(&n)

	drawer = make(map[int]bool)

	socks, _ := intScanln(n)

	for i := 0; i < n; i++ {
		if drawer[socks[i]] {
			count++
		}

		drawer[socks[i]] = !drawer[socks[i]]
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
