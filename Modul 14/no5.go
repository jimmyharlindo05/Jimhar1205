package main

import (
	"fmt"
)

// Mendefinisikan konstanta nMax sesuai soal
const nMax = 7919

// Definisi struct Buku
type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

// Definisi array DaftarBuku
type DaftarBuku [nMax]Buku

// Fungsi untuk mengurutkan buku berdasarkan rating secara menurun (Descending) dengan Insertion Sort
func insertionSortDescending(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1

		// Geser elemen yang ratingnya lebih kecil dari key.rating ke kanan
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func main() {
	var pustaka DaftarBuku
	var n, searchRating int

	// Membaca banyaknya data buku (N)
	fmt.Scan(&n)

	// Membaca data untuk masing-masing buku
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}

	// Membaca rating buku yang ingin dicari
	fmt.Scan(&searchRating)

	// Mengurutkan buku berdasarkan rating tertinggi (Descending)
	if n > 0 {
		insertionSortDescending(&pustaka, n)
	}

	// 1. Baris Pertama: Data buku terfavorit (Rating tertinggi, berada di indeks ke-0)
	if n > 0 {
		fav := pustaka[0]
		fmt.Printf("%s %s %s %s %d %d %d\n", fav.id, fav.judul, fav.penulis, fav.penerbit, fav.eksemplar, fav.tahun, fav.rating)
	}

	// 2. Baris Kedua: Lima judul buku dengan rating tertinggi
	limit := 5
	if n < limit {
		limit = n // Berjaga-jaga jika jumlah total buku kurang dari 5
	}
	
	for i := 0; i < limit; i++ {
		if i == limit-1 {
			fmt.Printf("%s\n", pustaka[i].judul) // Buku ke-5 di-print dengan newline
		} else {
			fmt.Printf("%s ", pustaka[i].judul)
		}
	}

	// 3. Baris Terakhir: Data buku yang dicari sesuai rating
	found := false
	for i := 0; i < n; i++ {
		// Menampilkan buku pertama yang ratingnya cocok dengan pencarian
		if pustaka[i].rating == searchRating {
			b := pustaka[i]
			fmt.Printf("%s %s %s %s %d %d %d\n", b.id, b.judul, b.penulis, b.penerbit, b.eksemplar, b.tahun, b.rating)
			found = true
			break 
		}
	}
	
	// Jika tidak ada buku dengan rating tersebut (Opsional untuk error handling)
	if !found && n > 0 {
		fmt.Println("Buku dengan rating tersebut tidak ditemukan.")
	}
}