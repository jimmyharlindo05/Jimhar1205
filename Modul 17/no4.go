package main

import "fmt"

func main() {
	var N int

	fmt.Print("N suku pertama: ")
	fmt.Scan(&N)

	var S float64

	for i := 0; i < N; i++ {
		suku := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			S += suku
		} else {
			S -= suku
		}
	}

	fmt.Printf("Hasil PI: %.7f\n", 4*S)

	var jumlah float64
	var i int
	var selisih float64 = 1

	for selisih >= 0.00001 {
		suku := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			jumlah += suku
		} else {
			jumlah -= suku
		}

		selisih = 1.0 / float64(2*(i+1)+1)
		i++
	}

	fmt.Printf("Hasil PI: %.9f\n", 4*jumlah)
	fmt.Printf("Pada i ke: %d\n", i)
}