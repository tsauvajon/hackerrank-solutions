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
		n, max, count int
	)

	scanner := bufio.NewReader(os.Stdin)

	fmt.Scan(&n)

	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values := strings.Split(buffer, " ")

	var integers = []int{}

	for _, num := range values {
		intValue, _ := strconv.Atoi(num)
		integers = append(integers, intValue)
	}

	for _, height := range integers {
		if height == max {
			count++
		}

		if height > max {
			max = height
			count = 1
		}
	}

	fmt.Print(count)
}
