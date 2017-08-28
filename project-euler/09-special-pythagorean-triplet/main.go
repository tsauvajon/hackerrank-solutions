package main

import "fmt"

func main() {
	var t, n int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)

		fmt.Println(specialPythagoreanTriplet(n))
	}
}

func specialPythagoreanTriplet(sum int) int64 {
	max := int64(-1)

	sum64 := int64(sum)

	for a := int64(0); a <= sum64/3; a++ {
		for b := a; b <= sum64/3; b++ {
			c := sum64 - (a + b)

			product := a * b * c

			if a*a+b*b == c*c && max < product {
				max = product
			}
		}
	}

	return max
}
