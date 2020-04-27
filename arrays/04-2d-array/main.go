package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// hourglassIndices returns all hourglass point coordinates
func hourglassIndices(i, j int32) [][]int32 {
	return [][]int32{
		[]int32{i, j},
		[]int32{i + 1, j},
		[]int32{i + 2, j},
		[]int32{i + 1, j + 1},
		[]int32{i, j + 2},
		[]int32{i + 1, j + 2},
		[]int32{i + 2, j + 2},
	}
}

func hourglassVal(arr [][]int32, i, j int32) int32 {
	val := int32(0)
	for _, h := range hourglassIndices(i, j) {
		val += arr[h[1]][h[0]]
	}
	return val
}

func hourglassSum(arr [][]int32) int32 {
	max := int32(math.MinInt32)
	for i := int32(0); i < 4; i++ {
		for j := int32(0); j < 4; j++ {
			val := hourglassVal(arr, i, j)
			if val > max {
				max = val
			}
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
