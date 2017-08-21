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
		n, m, count int
		a           = []int{}
		b           = []int{}
	)

	scanner := bufio.NewReader(os.Stdin)

	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values := strings.Split(buffer, " ")

	n, _ = strconv.Atoi(values[0])
	m, _ = strconv.Atoi(values[1])

	buffer, _ = scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values = strings.Split(buffer, " ")

	for _, num := range values {
		intValue, _ := strconv.Atoi(num)
		a = append(a, intValue)
	}

	buffer, _ = scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values = strings.Split(buffer, " ")

	for _, num := range values {
		intValue, _ := strconv.Atoi(num)
		b = append(b, intValue)
	}

	for i := a[len(a)-1]; i < b[0]; i++ {
		isBetween := true
		for j := 0; j < n && isBetween; j++ {
			if i%a[j] != 0 {
				isBetween = false
			}
		}

		for j := 0; j < m && isBetween; j++ {
			if b[j]%i != 0 {
				isBetween = false
			}
		}

		if isBetween {
			count++
		}
	}

	fmt.Print(count)
}
