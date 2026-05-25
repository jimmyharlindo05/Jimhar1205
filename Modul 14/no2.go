package main

import (
	"fmt"
)

// Fungsi Selection Sort untuk mengurutkan menaik (Ascending) - Ganjil
func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi Selection Sort untuk mengurutkan menurun (Descending) - Genap
func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	// Membaca banyaknya daerah
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		// Membaca banyaknya rumah di daerah tersebut
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		// Membaca nomor rumah dan memisahkannya ke array ganjil/genap
		for j := 0; j < m; j++ {
			var num int
			fmt.Scan(&num)
			if num%2 != 0 {
				ganjil = append(ganjil, num)
			} else {
				genap = append(genap, num)
			}
		}

		// Mengurutkan sesuai aturan soal
		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		// Menggabungkan array ganjil (depan) dan genap (belakang)
		hasil := append(ganjil, genap...)

		// Mencetak hasil
		for j := 0; j < len(hasil); j++ {
			if j == len(hasil)-1 {
				fmt.Printf("%d", hasil[j])
			} else {
				fmt.Printf("%d ", hasil[j])
			}
		}
		fmt.Println() // Pindah ke baris baru untuk daerah selanjutnya
	}
}