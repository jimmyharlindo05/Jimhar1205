package main

import (
	"fmt"
)

func main() {
	var angka int
	totalMasuk := 0
	totalValid := 0

	// Array untuk menyimpan jumlah suara calon 1-20
	var suara [21]int

	for {
		fmt.Scan(&angka)

		// Data berhenti jika input 0
		if angka == 0 {
			break
		}

		totalMasuk++

		// Validasi suara valid antara 1 - 20
		if angka >= 1 && angka <= 20 {
			suara[angka]++
			totalValid++
		}
	}

	// Output
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalValid)

	// Menampilkan calon yang mendapat suara
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}