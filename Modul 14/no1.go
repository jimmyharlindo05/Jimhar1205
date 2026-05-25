package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan algoritma Selection Sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari indeks dengan nilai minimum di sisa array yang belum terurut
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil yang ditemukan dengan elemen di indeks i
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	// Membaca banyaknya daerah (n)
	fmt.Scan(&n)

	// Loop untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		// Membaca banyaknya rumah di daerah tersebut (m)
		fmt.Scan(&m)

		// Membuat slice untuk menyimpan nomor rumah
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Urutkan nomor rumah menggunakan Selection Sort
		selectionSort(rumah)

		// Cetak hasil pengurutan untuk daerah saat ini
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Printf("%d", rumah[j]) // Elemen terakhir tanpa spasi di ujung
			} else {
				fmt.Printf("%d ", rumah[j])
			}
		}
		fmt.Println() // Pindah baris untuk daerah berikutnya
	}
}