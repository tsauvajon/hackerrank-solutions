package main

import "fmt"

func main() {
	var (
		actualDay, actualMonth, actualYear       int
		expectedDay, expectedMonth, expectedYear int
	)

	fmt.Scan(&actualDay)
	fmt.Scan(&actualMonth)
	fmt.Scan(&actualYear)

	fmt.Scan(&expectedDay)
	fmt.Scan(&expectedMonth)
	fmt.Scan(&expectedYear)

	switch {
	case actualYear < expectedYear:
		fmt.Println(0)

	case actualYear == expectedYear && actualMonth < expectedMonth:
		fmt.Println(0)

	case actualYear == expectedYear && actualMonth == expectedMonth && actualDay <= expectedDay:
		fmt.Println(0)

	case actualYear == expectedYear && actualMonth == expectedMonth:
		fmt.Println(15 * (actualDay - expectedDay))

	case actualYear == expectedYear:
		fmt.Println(500 * (actualMonth - expectedMonth))

	default:
		fmt.Println(10000)
	}
}
