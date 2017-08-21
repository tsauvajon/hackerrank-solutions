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
		s, t, a, b, m, n, applesOnHouse, orangesOnHouse int
		apples                                          = []int{}
		oranges                                         = []int{}
	)

	scanner := bufio.NewReader(os.Stdin)

	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values := strings.Split(buffer, " ")

	s, _ = strconv.Atoi(values[0])
	t, _ = strconv.Atoi(values[1])

	buffer, _ = scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values = strings.Split(buffer, " ")

	a, _ = strconv.Atoi(values[0])
	b, _ = strconv.Atoi(values[1])

	buffer, _ = scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values = strings.Split(buffer, " ")

	m, _ = strconv.Atoi(values[0])
	n, _ = strconv.Atoi(values[1])

	buffer, _ = scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values = strings.Split(buffer, " ")

	for _, num := range values {
		intValue, _ := strconv.Atoi(num)
		apples = append(apples, intValue)
	}

	buffer, _ = scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values = strings.Split(buffer, " ")

	for _, num := range values {
		intValue, _ := strconv.Atoi(num)
		oranges = append(oranges, intValue)
	}

	for i := 0; i < m; i++ {
		if a+apples[i] >= s && a+apples[i] <= t {
			applesOnHouse++
		}
	}

	for i := 0; i < n; i++ {
		if b+oranges[i] >= s && b+oranges[i] <= t {
			orangesOnHouse++
		}
	}

	fmt.Printf("%d\n%d", applesOnHouse, orangesOnHouse)
}
