package main

import "fmt"

func main() {
	var N int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&N)

	var jumlah float64

	for i := 0; i < N; i++ {
		suku := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			jumlah += suku
		} else {
			jumlah -= suku
		}
	}

	fmt.Printf("Hasil PI: %.10f\n", 4*jumlah)

	jumlah = 0
	var piLama, piBaru float64
	i := 0

	for {
		suku := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			jumlah += suku
		} else {
			jumlah -= suku
		}

		piLama = piBaru
		piBaru = 4 * jumlah

		if i > 0 {
			selisih := piBaru - piLama
			if selisih < 0 {
				selisih = -selisih
			}

			if selisih <= 0.00001 {
				break
			}
		}

		i++
	}

	fmt.Printf("Hasil PI: %.10f\n", piLama)
	fmt.Printf("Hasil PI: %.10f\n", piBaru)
	fmt.Printf("Pada i ke: %d\n", i+1)
}