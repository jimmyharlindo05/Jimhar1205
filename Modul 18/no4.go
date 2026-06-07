package main

import (
	"fmt"
)

var pita string
var idx int
var karakter byte

func start() {
	idx = 0
	karakter = pita[idx]
}

func maju() {
	idx++
	if idx < len(pita) {
		karakter = pita[idx]
	}
}

func eop() bool {
	return karakter == '.'
}

func cc() byte {
	return karakter
}

func main() {
	var jumlahKarakter, jumlahA, jumlahLE int
	var prev byte

	fmt.Print("Masukkan rangkaian karakter (akhiri dengan .): ")
	fmt.Scan(&pita)

	start()

	for !eop() {

		fmt.Print(string(cc()))

		jumlahKarakter++

		if cc() == 'A' {
			jumlahA++
		}

		if prev == 'L' && cc() == 'E' {
			jumlahLE++
		}

		prev = cc()
		maju()
	}

	fmt.Println()
	fmt.Println("Jumlah karakter :", jumlahKarakter)
	fmt.Println("Jumlah huruf A  :", jumlahA)

	if jumlahKarakter > 0 {
		fmt.Printf("Frekuensi A     : %.2f\n",
			float64(jumlahA)/float64(jumlahKarakter))
	} else {
		fmt.Println("Frekuensi A     : 0")
	}

	fmt.Println("Jumlah kata LE  :", jumlahLE)
}