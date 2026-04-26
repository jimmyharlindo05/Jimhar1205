package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ikan [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var wadah [1000]float64
	index := 0
	jumlahWadah := 0

	for i := 0; i < x; i++ {
		wadah[index] += ikan[i]

		// pindah ke wadah berikutnya setiap y ikan
		if (i+1)%y == 0 {
			index++
			jumlahWadah++
		}
	}

	// jika sisa ikan tidak habis dibagi y
	if x%y != 0 {
		jumlahWadah++
	}

	// output total tiap wadah
	totalSemua := 0.0
	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(wadah[i], " ")
		totalSemua += wadah[i]
	}

	fmt.Println()

	// rata-rata
	rata := totalSemua / float64(jumlahWadah)
	fmt.Println(rata)
}