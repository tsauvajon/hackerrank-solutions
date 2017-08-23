package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int

	scanner := bufio.NewReader(os.Stdin)

	fmt.Scan(&n)

	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	numbers := strings.Split(buffer, " ")

	for i := n - 1; i >= 0; i-- {
		fmt.Print(numbers[i], " ")
	}
}
