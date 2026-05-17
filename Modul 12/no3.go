package main

import "fmt"

const NMAX = 100000

var data [NMAX]int

func main() {
	var n, k int

	// Input jumlah data dan angka yang dicari
	fmt.Scan(&n, &k)

	// Mengisi array
	isiArray(n)

	// Mencari posisi
	hasil := posisi(n, k)

	// Output
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

// Procedure untuk mengisi array
func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

// Function mencari posisi k dalam array
func posisi(n, k int) int {
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i // posisi dimulai dari 0
		}
	}
	return -1
}