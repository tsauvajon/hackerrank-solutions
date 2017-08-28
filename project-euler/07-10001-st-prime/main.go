package main

import "fmt"

func main() {
	var (
		t int
	)

	fmt.Scan(&t)

	// 10000th prime : 104729
	primes := sieveOfEratosthene(104730)

	for i := 0; i < t; i++ {
		var input int
		fmt.Scan(&input)

		fmt.Println(primes[input-1])
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

/* v1
func main() {
	var (
		t         int
		testCases []int
	)

	fmt.Scan(&t)

	testCases = make([]int, t)

	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, n := range testCases {
		fmt.Println(nthPrime(n))
	}
}

func nthPrime(n int) uint {
	var m uint

	for i := 0; i < n; i++ {
		m = nextPrime(m)
	}

	return m
}

func nextPrime(n uint) uint {
	n++

	for largestPrimeFactor(n) != n {
		n++
	}

	return n
}

func largestPrimeFactor(n uint) uint {
	var (
		i, max uint
	)

	i = 2
	max = 2

	for i*i <= n {
		for n%i == 0 {
			max = i
			n /= i
		}

		i++
	}

	if n > max {
		max = n
	}

	return max
}
*/

/* v2

func main() {
	var (
		t         int
		testCases []int
	)

	fmt.Scan(&t)

	testCases = make([]int, t, t)

	for i := range testCases {
		fmt.Scan(&testCases[i])
	}

	for _, n := range testCases {
		fmt.Println(nthPrime(n))
	}
}

func nthPrime(n int) uint {
	var (
		end     int
		current uint
		stack   []uint
	)

	current = 5
	end = 2

	stack = make([]uint, n)
	stack[0] = 2
	stack[1] = 3

	for n > end {
		if isPrime(end, current, stack) {
			stack[end] = current
			end++
		}

		current += 2
	}

	return stack[n-1]
}

func isPrime(end int, curr uint, stack []uint) bool {
	for i := 0; i < end; i++ {
		if curr%stack[i] == 0 {
			return false
		}
	}

	return true
}

*/
