package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	
 	
    // Declare second integer, double, and String variables.
    var (
        j uint64
        e float64
        t string
    )
    
    // Read and save an integer, double, and String to your variables.
    fmt.Scan(&j)
    fmt.Scan(&e)
    scanner.Scan()
    t = scanner.Text()
    
    // Print the sum of both integer variables on a new line.
    fmt.Println(i + j)

    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f\n", d + e)
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Println(s + t)
}