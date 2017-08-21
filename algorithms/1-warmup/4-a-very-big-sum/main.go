package main

import (
    "fmt"
)

func main() {
    var n int;
    fmt.Scan(&n);
    
    input := make([]int64, n)
    
    for i := range input {
        fmt.Scan(&input[i])
    }
    
    var total int64
    
    for i := range input {
        total += input[i]
    }
    
    fmt.Print(total)
}