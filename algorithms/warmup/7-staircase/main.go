package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	for i := 1; i < n+1; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {
			fmt.Print("#")
		}

		fmt.Println()
	}
}
