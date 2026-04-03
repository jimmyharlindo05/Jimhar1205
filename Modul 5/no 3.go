package main

import "fmt"

// fungsi rekursif untuk mencari faktor
func faktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor: ")
	faktor(n, 1)
}