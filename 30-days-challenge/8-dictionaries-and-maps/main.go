package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		n int
		m map[string]string
	)

	scanner := bufio.NewReader(os.Stdin)
	m = make(map[string]string)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		buffer, _ := scanner.ReadString('\n')
		buffer = strings.Trim(buffer, "\r\n")
		kvp := strings.Split(buffer, " ")
		m[kvp[0]] = kvp[1]
	}

	for {
		buffer, _ := scanner.ReadString('\n')

		if buffer == "" {
			break
		}

		buffer = strings.Trim(buffer, "\r\n")

		if m[buffer] == "" {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", buffer, m[buffer])
		}
	}
}
