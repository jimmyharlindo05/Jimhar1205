package main

import "fmt"

// fungsi rekursif
func barisan(n int) {
	if n == 0 {
		return
	}

	// cetak saat turun
	fmt.Print(n, " ")

	barisan(n - 1)

	// cetak saat naik (kecuali n=1 agar tidak dobel tengah)
	if n != 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	barisan(n)
}