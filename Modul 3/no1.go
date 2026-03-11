package main

import (
	"fmt"
)

// fungsi factorial
func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// fungsi permutation P(n,r)
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// fungsi combination C(n,r)
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	// input
	fmt.Scan(&a, &b, &c, &d)

	// baris pertama
	fmt.Println(permutation(a, c), combination(a, c))

	// baris kedua
	fmt.Println(permutation(b, d), combination(b, d))
}