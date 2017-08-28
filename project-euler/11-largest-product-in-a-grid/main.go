package main

import "fmt"

var (
	grid [][]int
)

func main() {
	grid = make([][]int, 20)

	for i := 0; i < 20; i++ {
		grid[i] = make([]int, 20)

		for j := 0; j < 20; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	product := 0

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			value := max(right(j, i), bottom(j, i), rightDiag(j, i), leftDiag(j, i))
			if value > product {
				product = value
			}
		}
	}

	fmt.Println(product)
}

func right(i, j int) int {
	if j > 16 {
		return 0
	}

	product := 1

	for k := 0; k < 4; k++ {
		product *= grid[i][j+k]
	}

	return product
}

func bottom(i, j int) int {
	if i > 16 {
		return 0
	}

	product := 1

	for k := 0; k < 4; k++ {
		product *= grid[i+k][j]
	}

	return product
}

func rightDiag(i, j int) int {
	if j > 16 || i > 16 {
		return 0
	}

	product := 1

	for k := 0; k < 4; k++ {
		product *= grid[i+k][j+k]
	}

	return product
}

func leftDiag(i, j int) int {
	if i < 3 || j > 16 {
		return 0
	}

	product := 1

	for k := 0; k < 4; k++ {
		product *= grid[i-k][j+k]
	}

	return product
}

func max(a, b, c, d int) int {
	max := a

	if b > max {
		max = b
	}

	if c > max {
		max = c
	}

	if d > max {
		return d
	}

	return max
}
