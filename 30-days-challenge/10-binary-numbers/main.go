package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		input, max, current int64
	)

	fmt.Scan(&input)

	binary := strconv.FormatInt(input, 2)

	for i := 0; i < len(binary); i++ {
		if string(binary[i]) == "0" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current++
		}
	}

	if current > max {
		max = current
	}

	fmt.Print(max)
}
