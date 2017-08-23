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
		min, max int64
	)

	scanner := bufio.NewReader(os.Stdin)

	var totals = []int64{}

	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values := strings.Split(buffer, " ")

	var integers = []int64{}

	for _, num := range values {
		intValue, _ := strconv.Atoi(num)
		integers = append(integers, int64(intValue))
	}

	for i := 0; i < len(integers); i++ {
		current := int64(0)

		for j := 0; j < len(integers); j++ {
			if i == j {
				continue
			}

			current += int64(integers[j])
		}

		totals = append(totals, current)
	}

	max = totals[0]
	min = totals[0]

	for i := 1; i < len(totals); i++ {
		if totals[i] > max {
			max = totals[i]
		}

		if totals[i] < min {
			min = totals[i]
		}
	}

	fmt.Printf("%d %d", min, max)
}
