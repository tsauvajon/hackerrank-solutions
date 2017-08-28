package main

import "fmt"

func main() {
	var (
		t int
	)

	fmt.Scan(&t)

	// 10^6th prime = 15485863
	primes := sieveOfEratosthene(15485864)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		fmt.Println(summation(n, primes))
	}
}

func sieveOfEratosthene(n int) (primes []int) {
	b := make([]bool, n)

	for i := 2; i < n; i++ {
		if b[i] {
			continue
		}

		primes = append(primes, i)

		for k := i * i; k < n; k += i {
			b[k] = true
		}
	}
	return
}

func summation(limit int, primes []int) int {
	sum := 0
	for _, prime := range primes {
		if prime > limit {
			return sum
		}
		sum += prime
	}

	return sum
}
