package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n, primaryDiagonal, secondaryDiagonal int
	)

	scanner := bufio.NewReader(os.Stdin)

	fmt.Scan(&n)

	m := make([][]int, n)

	for i := 0; i < n; i++ {
		buffer, _ := scanner.ReadString('\n')
		buffer = strings.Trim(buffer, "\r\n")
		values := strings.Split(buffer, " ")
		var integers = []int{}

		for _, num := range values {
			intValue, _ := strconv.Atoi(num)
			integers = append(integers, intValue)
		}

		m[i] = integers
	}

	for i := 0; i < n; i++ {
		primaryDiagonal += m[i][i]
		secondaryDiagonal += m[i][n-1-i]
	}

	sum := primaryDiagonal - secondaryDiagonal

	if sum < 0 {
		sum = -sum
	}

	fmt.Print(sum)
}
