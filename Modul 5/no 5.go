package main

import "fmt"

// fungsi rekursif
func ganjil(i int, n int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Print(i, " ")
	}

	ganjil(i+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	ganjil(1, n)
}