package main

import (
	"fmt"
	"math"
)

// tipe data titik
type titik struct {
	x, y float64
}

// tipe data lingkaran
type lingkaran struct {
	pusat titik
	r     float64
}

// fungsi menghitung jarak dua titik
func jarak(p, q titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

// fungsi cek apakah titik di dalam lingkaran
func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= c.r
}

func main() {
	var c1, c2 lingkaran
	var p titik

	// input lingkaran 1
	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.r)

	// input lingkaran 2
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.r)

	// input titik
	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	// output sesuai kondisi
	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}