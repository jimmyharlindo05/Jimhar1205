package main

import "fmt"

// prosedur hitungSkor
func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		// jika <= 300 berarti berhasil
		if waktu <= 300 {
			*soal = *soal + 1
			*skor = *skor + waktu
		}
	}
}

func main() {
	var nama string
	var pemenang string
	maxSoal := -1
	minSkor := 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		// tentukan pemenang
		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}

	// output
	fmt.Println(pemenang, maxSoal, minSkor)
}