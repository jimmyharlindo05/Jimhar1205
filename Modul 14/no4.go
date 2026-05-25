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

		// Geser elemen yang lebih besar ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int
	var num int

	// Membaca masukan hingga menemukan bilangan negatif
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break // Berhenti jika tidak ada lagi input
		}
		
		if num < 0 {
			break // Angka negatif menandakan akhir dari input
		}
		
		// Menyimpan hanya bilangan non-negatif
		arr = append(arr, num)
	}

	// Memastikan array tidak kosong sebelum diproses
	if len(arr) > 0 {
		// Mengurutkan data
		insertionSort(arr)

		// Mencetak baris pertama: isi array setelah diurutkan
		for i, val := range arr {
			if i == len(arr)-1 {
				fmt.Printf("%d\n", val)
			} else {
				fmt.Printf("%d ", val)
			}
		}

		// Memeriksa jarak/selisih antar elemen
		if len(arr) < 2 {
			fmt.Println("Data berjarak tidak tetap")
		} else {
			jarakAwal := arr[1] - arr[0]
			isTetap := true

			// Mengecek apakah semua jarak sama dengan jarakAwal
			for i := 2; i < len(arr); i++ {
				if arr[i]-arr[i-1] != jarakAwal {
					isTetap = false
					break
				}
			}

			// Mencetak baris kedua: status jarak
			if isTetap {
				fmt.Printf("Data berjarak %d\n", jarakAwal)
			} else {
				fmt.Println("Data berjarak tidak tetap")
			}
		}
	}
}