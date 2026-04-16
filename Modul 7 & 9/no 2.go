package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	// input array
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. tampilkan semua isi array
	fmt.Println("Isi array:", arr)

	// b. indeks ganjil
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. indeks genap
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. kelipatan x
	var x int
	fmt.Print("Masukkan kelipatan x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. hapus elemen indeks tertentu
	var idx int
	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Array setelah dihapus:", arr)

	// f. rata-rata
	total := 0
	for _, v := range arr {
		total += v
	}
	rata := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata:", rata)

	// g. standar deviasi
	var jumlah float64
	for _, v := range arr {
		jumlah += math.Pow(float64(v)-rata, 2)
	}
	std := math.Sqrt(jumlah / float64(len(arr)))
	fmt.Println("Standar deviasi:", std)

	// h. frekuensi angka tertentu
	var cari, freq int
	fmt.Print("Masukkan angka yang dicari: ")
	fmt.Scan(&cari)

	for _, v := range arr {
		if v == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}