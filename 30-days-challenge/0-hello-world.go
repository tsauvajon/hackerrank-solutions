package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
	fmt.Println("Hello, World.")

	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('\n')
    
	fmt.Println(text)
}