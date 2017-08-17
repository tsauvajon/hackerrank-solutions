package main

import "fmt"

func main() {
    var (
        meal, tip, tax float64
    )
    
    fmt.Scan(&meal)
    fmt.Scan(&tip)
    fmt.Scan(&tax)
    
    total := meal * (1 + (tip / 100) + (tax / 100))
    fmt.Printf("The total meal cost is %d dollars.", int(total + .5))
}