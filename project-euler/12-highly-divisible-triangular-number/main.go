package main

import "fmt"

func main() {
	var t, n int

	m := generateMap()

	fmt.Scan(&t)

	fmt.Println(getNbDivisorsv2(6))
	fmt.Println(getNbDivisorsv2(15))
	fmt.Println(getNbDivisorsv2(28))

	for i := 0; i < t; i++ {
		fmt.Scan(&n)

		for m[n] == 0 {
			n++
		}

		fmt.Println(m[n])
	}
}

func generateMap() map[int]int {
	triangles := make(map[int]int)

	triangle := 1

	triangles[1] = 1

	for i := 2; i < 10000; i++ {
		triangle += i

		nbDivisors := getNbDivisorsv2(triangle)

		if triangles[nbDivisors] == 0 {
			triangles[nbDivisors] = triangle
		}
	}

	return triangles
}

func getNbDivisors(n int) int {
	nbDivisors := 2

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			nbDivisors++
		}
	}

	return nbDivisors
}

func getNbDivisorsv2(n int) int {
	if n == 1 {
		return 1
	}

	limit := n
	nb := 0

	for i := 2; i < limit; i++ {
		if n%i == 0 {
			limit = n / i

			if limit != i {
				nb++
			}
		}
		nb++
	}

	return nb
}
