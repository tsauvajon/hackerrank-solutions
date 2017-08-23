package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n     int
		words []string
	)

	scanner := bufio.NewReader(os.Stdin)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		word, _ := scanner.ReadString('\n')
		words = append(words, word)
	}

	for _, word := range words {
		for i := 0; i < len(word); i += 2 {
			if word[i] != '\n' {
				fmt.Printf("%c", word[i])
			}
		}

		fmt.Printf(" ")

		for i := 1; i < len(word); i += 2 {
			if word[i] != '\n' {
				fmt.Printf("%c", word[i])
			}
		}

		fmt.Println()
	}
}
