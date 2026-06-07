package main

import "fmt"

func main() {
	var x string
	var n int

	fmt.Print("Masukkan string yang dicari: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah data string: ")
	fmt.Scan(&n)

	data := make([]string, n)

	fmt.Println("Masukkan data string:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	ditemukan := false
	posisi := -1
	jumlah := 0

	for i := 0; i < n; i++ {
		if data[i] == x {
			if !ditemukan {
				posisi = i + 1
				ditemukan = true
			}
			jumlah++
		}
	}
	
	if ditemukan {
		fmt.Println("a. String ditemukan")
	} else {
		fmt.Println("a. String tidak ditemukan")
	}

	if ditemukan {
		fmt.Println("b. Posisi pertama:", posisi)
	} else {
		fmt.Println("b. Tidak memiliki posisi")
	}

	fmt.Println("c. Jumlah kemunculan:", jumlah)

	if jumlah >= 2 {
		fmt.Println("d. Ya, muncul sedikitnya dua kali")
	} else {
		fmt.Println("d. Tidak, muncul kurang dari dua kali")
	}
}