package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	var hasil []string
	i := 1

	for {
		fmt.Print("Pertandingan ", i, ": ")
		fmt.Scan(&skorA, &skorB)

		// berhenti jika ada skor negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			fmt.Println("Hasil", i, ":", klubA)
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			fmt.Println("Hasil", i, ":", klubB)
			hasil = append(hasil, klubB)
		} else {
			fmt.Println("Hasil", i, ": Draw")
		}

		i++
	}

	fmt.Println("Pertandingan selesai")

	// tampilkan klub yang menang
	fmt.Println("Daftar pemenang:")
	for i, v := range hasil {
		fmt.Println("Menang", i+1, ":", v)
	}
}