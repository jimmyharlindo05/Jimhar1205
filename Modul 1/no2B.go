package main

import "fmt"

func main() {
	var N int
	var bunga string
	pita := ""

	fmt.Print("N : ")
	fmt.Scan(&N)

	if N == 0 {
		fmt.Println("Pita :")
		return
	}

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d : ", i)
		fmt.Scan(&bunga)

		pita = pita + bunga + " - "
	}

	fmt.Println("Pita :", pita)
}