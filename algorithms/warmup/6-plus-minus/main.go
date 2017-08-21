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
		n, negatives, zeroes, positives int
		nf, zf, pf                      float64
	)

	scanner := bufio.NewReader(os.Stdin)

	fmt.Scan(&n)

	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values := strings.Split(buffer, " ")

	for _, num := range values {
		intValue, _ := strconv.Atoi(num)
		switch {
		case intValue < 0:
			negatives++
		case intValue == 0:
			zeroes++
		default:
			positives++
		}
	}

	nf = float64(negatives) / float64(n)
	zf = float64(zeroes) / float64(n)
	pf = float64(positives) / float64(n)

	fmt.Printf("%.6f\n%.6f\n%.6f", pf, nf, zf)
}
