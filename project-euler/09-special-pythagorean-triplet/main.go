package main

import "fmt"
import "math/big"

func main() {
	var (
		t int
		n int64
	)

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)

		fmt.Println(specialPythagoreanTriplet(n))
	}
}

func specialPythagoreanTriplet(sum int64) int64 {
	// max := int64(-1)
	max := big.NewInt(-1)

	for a := int64(1); a <= sum/3; a++ {
		b := (sum*sum - 2*sum*a) / (2*sum - 2*a)
		c := sum - (a + b)

		aSquare := big.NewInt(0).Mul(big.NewInt(a), big.NewInt(a))
		bSquare := big.NewInt(0).Mul(big.NewInt(b), big.NewInt(b))
		cSquare := big.NewInt(0).Mul(big.NewInt(c), big.NewInt(c))

		abSquareSum := big.NewInt(0).Add(aSquare, bSquare)

		if abSquareSum.Cmp(cSquare) != 0 {
			continue
		}

		product := big.NewInt(0).Mul(big.NewInt(a), big.NewInt(b))
		product = product.Mul(product, big.NewInt(c))
		if max.Cmp(product) == -1 {
			max = product
		}
	}

	return max.Int64()
}
