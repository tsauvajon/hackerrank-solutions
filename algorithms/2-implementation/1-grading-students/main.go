package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)

		diff := 5 - (val % 5)

		switch {
		case val < 38:
			fmt.Println(val)
		case diff < 3:
			fmt.Println(val + diff)
		default:
			fmt.Println(val)
		}
	}
}
