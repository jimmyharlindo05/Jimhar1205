package main

import (
	"fmt"
)

func main() {
	var angka int
	totalMasuk := 0
	totalValid := 0

	// Menyimpan jumlah suara calon 1-20
	var suara [21]int

	// Input data
	for {
		fmt.Scan(&angka)

		// Berhenti jika 0
		if angka == 0 {
			break
		}

		totalMasuk++

		// Validasi suara sah
		if angka >= 1 && angka <= 20 {
			suara[angka]++
			totalValid++
		}
	}

	// Menentukan ketua dan wakil
	ketua := 1
	wakil := 1

	for i := 2; i <= 20; i++ {
		// Cari suara terbanyak untuk ketua
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if suara[i] > suara[wakil] && i != ketua {
			wakil = i
		}
	}

	// Pastikan wakil adalah suara terbanyak kedua
	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > suara[wakil] ||
				(suara[i] == suara[wakil] && i < wakil) {
				wakil = i
			}
		}
	}

	// Output
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalValid)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}