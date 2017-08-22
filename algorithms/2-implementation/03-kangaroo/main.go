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
		x1, x2, v1, v2           int
		theFasterKangarooIsAhead bool
	)

	scanner := bufio.NewReader(os.Stdin)

	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	values := strings.Split(buffer, " ")

	x1, _ = strconv.Atoi(values[0])
	v1, _ = strconv.Atoi(values[1])
	x2, _ = strconv.Atoi(values[2])
	v2, _ = strconv.Atoi(values[3])

	if v1 == v2 && x1 != x2 {
		fmt.Print("NO")
		return
	}

	theFasterKangarooIsAhead = isTheFasterKangarooAhead(x1, x2, v1, v2)

	for x1 != x2 && !theFasterKangarooIsAhead {
		x1 += v1
		x2 += v2

		theFasterKangarooIsAhead = isTheFasterKangarooAhead(x1, x2, v1, v2)
	}

	if x1 == x2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

func isTheFasterKangarooAhead(x1, x2, v1, v2 int) bool {
	if v1 > v2 && x1 > x2 {
		return true
	}

	if v2 > v1 && x2 > x1 {
		return true
	}

	return false
}
