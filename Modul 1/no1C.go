package main
import "fmt"

func main() {

	var berat int
	var kg, gram int
	var biayaKg, biayaGram, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	gram = berat % 1000

	fmt.Printf("Detail berat: %d kg + %d g\n", kg, gram)

	biayaKg = kg * 10000

	if gram >= 500 {
		biayaGram = gram * 5
	} else {
		biayaGram = gram * 15
	}

	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGram)

	total = biayaKg + biayaGram

	if kg > 10 {
		total = biayaKg
	}

	fmt.Printf("Total biaya: Rp. %d\n", total)
}