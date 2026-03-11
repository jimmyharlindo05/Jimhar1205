package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara titik (a,b) dan (c,d)
func jarak(a, b, c, d float64) float64 {
	// Menggunakan rumus jarak: sqrt((a-c)^2 + (b-d)^2)
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi untuk mengecek apakah titik (x,y) berada di dalam lingkaran (cx,cy) dengan radius r
func didalam(cx, cy, r, x, y float64) bool {
	// Titik dianggap di dalam lingkaran jika jarak titik ke pusat <= radius
	jarakTitik := jarak(cx, cy, x, y)
	return jarakTitik <= r
}

func main() {
	// Deklarasi variabel
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	// Menerima input baris pertama (Lingkaran 1)
	fmt.Scan(&cx1, &cy1, &r1)
	
	// Menerima input baris kedua (Lingkaran 2)
	fmt.Scan(&cx2, &cy2, &r2)
	
	// Menerima input baris ketiga (Titik sembarang)
	fmt.Scan(&x, &y)

	// Mengecek posisi titik terhadap kedua lingkaran
	diDalamL1 := didalam(cx1, cy1, r1, x, y)
	diDalamL2 := didalam(cx2, cy2, r2, x, y)

	// Menentukan output berdasarkan kondisi
	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}