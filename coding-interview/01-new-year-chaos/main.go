package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// O(n)
func minimumBribes(q []int32) {
	var swaps int32
	min := int32(len(q))

	for i := len(q) - 1; i >= 0; i-- {
		if diff := q[i] - int32(i+1); diff > 2 {
			// case when someone bribed more than 2 persons
			fmt.Println("Too chaotic")
			return
		} else if diff > 0 {
			// case when someone got ahead by 1 or 2 positions
			swaps += diff
		} else {
			// check swaps that didn't result in final position gain
			if q[i] > min {
				swaps++
			} else if min > q[i] {
				min = q[i]
			}
		}
	}

	fmt.Println(swaps)
}

// Time outs on large sets, O(n²)
func minimumBribesV1(q []int32) {
	var swaps uint32

	for i := 0; i < len(q)-1; i++ {
		if q[i]-int32(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j := i + 1; j < len(q); j++ {
			if q[i] > q[j] {
				swaps++
			}
		}
	}

	fmt.Println(swaps)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
