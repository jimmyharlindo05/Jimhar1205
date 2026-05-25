package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan algoritma Insertion Sort
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen-elemen yang lebih besar dari key ke posisi setelahnya
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var data []int
	var num int

	for {
		// Membaca input bilangan bulat satu per satu
		_, err := fmt.Scan(&num)
		if err != nil {
			break // Berhenti jika tidak ada lagi input yang bisa dibaca
		}

		// Jika menemukan marker -5313, hentikan program
		if num == -5313 {
			break
		}

		// Jika menemukan angka 0, cetak median dari data yang terkumpul
		if num == 0 {
			n := len(data)
			if n == 0 {
				continue // Abaikan jika belum ada data yang tersimpan
			}

			// Mengurutkan data yang telah tersimpan menggunakan Insertion Sort
			insertionSort(data)

			// Mencari dan mencetak median
			if n%2 != 0 {
				// Jika jumlah data ganjil, ambil nilai tengah langsung
				fmt.Println(data[n/2])
			} else {
				// Jika jumlah data genap, ambil rata-rata dari dua nilai tengah
				// Pembagian integer di Go otomatis membulatkan ke bawah
				median := (data[(n/2)-1] + data[n/2]) / 2
				fmt.Println(median)
			}
		} else {
			// Jika bukan 0 dan bukan -5313, masukkan angka ke dalam slice
			data = append(data, num)
		}
	}
}