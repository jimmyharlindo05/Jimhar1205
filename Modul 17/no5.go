package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var banyakTopping int
	var toppingPizza int
	var x, y float64

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&banyakTopping)

	rand.Seed(1)

	for i := 0; i < banyakTopping; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			toppingPizza++
		}
	}

	pi := 4.0 * float64(toppingPizza) / float64(banyakTopping)

	fmt.Println("Topping pada Pizza:", toppingPizza)
	fmt.Printf("PI : %.9f\n", pi)
}