package main

import "fmt"

// Tipe data kartu domino
type Domino struct {
	Suit1 int  // gambar/suit sisi pertama
	Suit2 int  // gambar/suit sisi kedua
	Balak bool // true jika kedua sisi sama
}

// Satu set kartu domino
type Dominoes struct {
	Kartu [28]Domino
	Sisa  int
}

// Prosedur membuat satu set domino
func kokokKartu(D *Dominoes) {
	idx := 0

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			D.Kartu[idx] = Domino{
				Suit1: i,
				Suit2: j,
				Balak: i == j,
			}
			idx++
		}
	}

	D.Sisa = idx
}

// Fungsi mengambil kartu teratas
func ambilKartu(D *Dominoes) Domino {
	if D.Sisa == 0 {
		return Domino{}
	}

	D.Sisa--
	return D.Kartu[D.Sisa]
}

// Fungsi mengambil gambar/suit kartu
// suit = 1 untuk sisi pertama, selain itu sisi kedua
func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.Suit1
	}
	return d.Suit2
}

// Fungsi menghitung nilai kartu
func nilaiKartu(d Domino) int {
	return d.Suit1 + d.Suit2
}

func main() {
	var setDomino Dominoes

	kokokKartu(&setDomino)

	fmt.Println("Jumlah kartu:", setDomino.Sisa)

	kartu := ambilKartu(&setDomino)

	fmt.Println("Kartu yang diambil:")
	fmt.Println("Sisi 1 =", gambarKartu(kartu, 1))
	fmt.Println("Sisi 2 =", gambarKartu(kartu, 2))
	fmt.Println("Balak   =", kartu.Balak)
	fmt.Println("Nilai   =", nilaiKartu(kartu))

	fmt.Println("Sisa kartu:", setDomino.Sisa)
}