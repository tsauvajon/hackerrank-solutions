package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		n, sum, number int64
	)

	fmt.Scan(&n)

	for i := int64(0); i < n; i++ {
		fmt.Scan(&number)

		// // v1 iteration
		// sum = int64(0)

		// for j := int64(3); j < number; j++ {
		// 	if j%3 == 0 || j%5 == 0 {
		// 		sum += j
		// 	}
		// }

		// v2 summation
		// max := number - 1

		// factors3 = int64(max / int64(3))
		// highestFactor3 = max - (max % int64(3))
		// sum3 := float64((factors3 * (highestFactor3 + int64(3))) >> 1)

		// factors5 = int64(max / int64(5))
		// highestFactor5 = max - (max % int64(5))
		// sum5 := float64((factors5 * (highestFactor5 + int64(5))) >> 1)

		// // gcd = 3 * 5 = 15
		// factors15 = int64(max / int64(15))
		// highestFactor15 = max - (max % int64(15))
		// sum15 := float64((factors15 * (highestFactor15 + int64(15))) >> 1)

		// sum = int64(sum3 + sum5 - sum15)

		// fmt.Println(sum)

		// v3 summation + big ints
		max := big.NewInt(number - 1)
		two := big.NewInt(2)
		three := big.NewInt(3)
		five := big.NewInt(5)
		fifteen := big.NewInt(15)

		factors3 := big.NewInt(0).Div(max, three)
		factors3 = big.NewInt(factors3.Int64())
		quotient3 := big.NewInt(0).Rem(max, three)
		highestFactor3 := big.NewInt(0).Sub(max, quotient3)
		sum3 := big.NewInt(0).Div(big.NewInt(0).Mul(big.NewInt(0).Add(highestFactor3, three), factors3), two)

		factors5 := big.NewInt(0).Div(max, five)
		factors5 = big.NewInt(factors5.Int64())
		quotient5 := big.NewInt(0).Rem(max, five)
		highestFactor5 := big.NewInt(0).Sub(max, quotient5)
		sum5 := big.NewInt(0).Div(big.NewInt(0).Mul(big.NewInt(0).Add(highestFactor5, five), factors5), two)

		factors15 := big.NewInt(0).Div(max, fifteen)
		factors15 = big.NewInt(factors15.Int64())
		quotient15 := big.NewInt(0).Rem(max, fifteen)
		highestFactor15 := big.NewInt(0).Sub(max, quotient15)
		sum15 := big.NewInt(0).Div(big.NewInt(0).Mul(big.NewInt(0).Add(highestFactor15, fifteen), factors15), two)

		sum = big.NewInt(0).Sub(big.NewInt(0).Add(sum3, sum5), sum15).Int64()

		fmt.Println(sum)
	}
}
