package main

import "fmt"

func main() {
    result := 0
    times := readInt()
    for i := 0; i < times; i++ {
        result += readInt()
    }
    
    fmt.Println(result)
}

func readInt() int {
    var n int
    _, err := fmt.Scanf("%d", &n)

    if err != nil {
        panic(err)
    }

    return n
}