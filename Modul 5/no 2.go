package main

import "fmt"

func printStars(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	printStars(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	printStars(n, 1)
}