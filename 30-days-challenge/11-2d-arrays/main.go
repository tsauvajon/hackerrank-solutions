package main

import "fmt"

func main() {
	var (
		matrix = [6][6]int{}
		sums   = [16]int{}
		max    int
	)

	for i := 0; i < 6; i++ {
		row, _ := intScanln(6)

		for j, val := range row {
			matrix[i][j] = val
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			hourglassNb := i*4 + j
			top := matrix[i][j] + matrix[i][j+1] + matrix[i][j+2]
			middle := matrix[i+1][j+1]
			bottom := matrix[i+2][j] + matrix[i+2][j+1] + matrix[i+2][j+2]
			sums[hourglassNb] = top + middle + bottom

			if hourglassNb == 0 || sums[hourglassNb] > max {
				max = sums[hourglassNb]
			}
		}
	}

	fmt.Print(max)
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
