package main

import "fmt"

// ======================
// TIPE DATA
// ======================

type Domino struct {
	Suit1 int
	Suit2 int
	Balak bool
}

type Dominoes struct {
	Kartu [28]Domino
	Sisa  int
}

// ======================
// OPERASI DASAR
// ======================

// Membuat satu set kartu domino
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

	D.Sisa = idx
}

// Mengambil satu kartu dari tumpukan
func ambilKartu(D *Dominoes) Domino {
	if D.Sisa == 0 {
		return Domino{}
	}

	D.Sisa--
	return D.Kartu[D.Sisa]
}

// Mengambil nilai gambar/suit kartu
func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.Suit1
	}
	return d.Suit2
}

// Menghitung nilai kartu
func nilaiKartu(d Domino) int {
	return d.Suit1 + d.Suit2
}

// ======================
// SOAL NOMOR 2
// ======================

// Mengecek apakah dua kartu memiliki suit yang sama
func cocok(d1, d2 Domino) bool {
	return d1.Suit1 == d2.Suit1 ||
		d1.Suit1 == d2.Suit2 ||
		d1.Suit2 == d2.Suit1 ||
		d1.Suit2 == d2.Suit2
}

// Menggali kartu sampai ditemukan kartu yang cocok
func galiKartu(D *Dominoes, acuan Domino) Domino {
	var kartu Domino

	for D.Sisa > 0 {
		kartu = ambilKartu(D)

		if cocok(kartu, acuan) {
			return kartu
		}
	}

	return Domino{}
}

// True jika total nilai dua kartu = 12
func sepasangKartu(d1, d2 Domino) bool {
	return nilaiKartu(d1)+nilaiKartu(d2) == 12
}

// ======================
// MAIN PROGRAM
// ======================

func main() {
	var setDomino Dominoes

	// Membuat satu set domino
	kocokKartu(&setDomino)

	fmt.Println("Jumlah kartu awal:", setDomino.Sisa)

	// Kartu acuan
	var acuan Domino
	acuan.Suit1 = 3
	acuan.Suit2 = 5
	acuan.Balak = false

	fmt.Println("\nKartu Acuan:")
	fmt.Println(acuan.Suit1, "|", acuan.Suit2)

	// Gali kartu
	hasil := galiKartu(&setDomino, acuan)

	fmt.Println("\nKartu Hasil Gali:")
	fmt.Println(hasil.Suit1, "|", hasil.Suit2)

	fmt.Println("Balak :", hasil.Balak)
	fmt.Println("Nilai :", nilaiKartu(hasil))

	// Cek sepasang kartu
	if sepasangKartu(acuan, hasil) {
		fmt.Println("\nSepasang Kartu = TRUE")
	} else {
		fmt.Println("\nSepasang Kartu = FALSE")
	}

	fmt.Println("Sisa kartu dalam tumpukan:", setDomino.Sisa)
}