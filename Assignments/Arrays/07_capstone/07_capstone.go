package main

import "fmt"

// Stage 1: Define the Book struct here

type Book struct {
	title string
	price float64
}

func main() {
    // Stage 2: Create the 'promos' map and add "FALL20" -> 5.00
	promos := map[string]float64{"FALL20": 5.00}

    // Stage 3: Create the 'cart' slice with "The Hobbit" (15.50) and "Dune" (22.00)
	cart := []Book{{title: "The Hobbit", price: 15.50}, {title: "Dune", price: 22.00}}

    // Stage 4: Loop over 'cart', adding each book's Price to the subtotal
    subtotal := 0.0
	for _, book := range cart {
		subtotal += book.price
	}


    // Stage 5: Subtract the "FALL20" discount from the promos map and print the final total
	finaltotal := subtotal - promos["FALL20"]
	fmt.Println("Final Total:", finaltotal)

}