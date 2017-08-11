package main

import "fmt"

func main() {
    arr := make([]int, 6)
    var (
        a int
        b int
    )
    
    for i := 0; i < 6; i++ {
        fmt.Scan(&arr[i])
    }
    
    for i := 0; i < 3; i ++ {
        if arr[i] > arr[3+i] {
            a++
        } else if arr[3+i] > arr[i] {
            b++
        }
    }
    
    fmt.Printf("%d %d", a, b)
}