package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

// isi array
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Scan(&input)

	*n = len(input)
	for i := 0; i < *n; i++ {
		t[i] = rune(input[i])
	}
}


// cetak array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// balik array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// cek palindrome
func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Teks: ")
	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// reverse
	balikanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	// cek palindrome (pakai array asli, jadi balik lagi)
	balikanArray(&tab, n)
	fmt.Print("Palindrom: ")
	if palindrome(tab, n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}