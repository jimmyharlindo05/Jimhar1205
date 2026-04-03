package main

import "fmt"

// Fungsi rekursif untuk menghitung pangkat
func pangkat(x int, y int) int {
	// Base case: jika pangkat adalah 0, kembalikan 1
	if y == 0 {
		return 1
	}
	// Rekursif: x dikali dengan fungsi pangkat itu sendiri (y berkurang 1)
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int

	// Meminta masukan dari pengguna
	fmt.Scan(&x, &y)

	// Memanggil fungsi dan mencetak hasil
	hasil := pangkat(x, y)
	fmt.Println(hasil)
}