package main

import "fmt"

func main() {
	var bil, jumlah, rata float64
	var banyak int

	fmt.Println("Masukkan bilangan real (akhiri dengan 9999):")
	fmt.Scan(&bil)

	for bil != 9999 {
		jumlah += bil
		banyak++
		fmt.Scan(&bil)
	}

	if banyak > 0 {
		rata = jumlah / float64(banyak)
		fmt.Printf("Rata-rata = %.2f\n", rata)
	} else {
		fmt.Println("Tidak ada data yang dimasukkan.")
	}
}