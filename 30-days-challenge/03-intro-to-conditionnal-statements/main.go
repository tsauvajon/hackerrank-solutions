package main

import "fmt"

func main() {
    var n int
    
    fmt.Scan(&n)
    
    switch {
        case n % 2 == 1 || (n >= 6 && n <= 20):
            fmt.Print("Weird")
        default:
            fmt.Print("Not Weird")
    }
}