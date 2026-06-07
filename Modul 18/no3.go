package main

import "fmt"

type Domino struct {
	Suit1 int
	Suit2 int
	Balak bool
}

type Dominoes struct {
	Kartu [28]Domino
	Sisa  int
}

func kocokKartu(D *Dominoes) {
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

	D.Sisa = 28
}

func ambilKartu(D *Dominoes) Domino {
	D.Sisa--
	return D.Kartu[D.Sisa]
}

func tampilKartu(d Domino) {
	fmt.Printf("[%d|%d]", d.Suit1, d.Suit2)
}

func bisaMain(k Domino, kiri, kanan int) bool {
	return k.Suit1 == kiri ||
		k.Suit2 == kiri ||
		k.Suit1 == kanan ||
		k.Suit2 == kanan
}

func main() {
	var deck Dominoes
	var tangan [7]Domino

	kocokKartu(&deck)

	for i := 0; i < 7; i++ {
		tangan[i] = ambilKartu(&deck)
	}

	rangkaian := tangan[0]
	kiri := rangkaian.Suit1
	kanan := rangkaian.Suit2

	fmt.Println("Kartu awal:")
	tampilKartu(rangkaian)
	fmt.Println()

	berhasil := 1

	for i := 1; i < 7; i++ {

		if bisaMain(tangan[i], kiri, kanan) {

			if tangan[i].Suit1 == kiri {
				kiri = tangan[i].Suit2
			} else if tangan[i].Suit2 == kiri {
				kiri = tangan[i].Suit1
			} else if tangan[i].Suit1 == kanan {
				kanan = tangan[i].Suit2
			} else if tangan[i].Suit2 == kanan {
				kanan = tangan[i].Suit1
			}

			fmt.Print("Mainkan kartu ")
			tampilKartu(tangan[i])
			fmt.Println()

			berhasil++
		}
	}

	fmt.Println()
	fmt.Println("Jumlah kartu berhasil dimainkan:", berhasil)

	if berhasil == 7 {
		fmt.Println("Pemain MENANG")
	} else {
		fmt.Println("Pemain belum menghabiskan semua kartu")
	}
}