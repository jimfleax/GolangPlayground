package main

import "fmt"

// Stage 1: Define the Recipe struct here
type Recipe struct {
	Name string
	Cost float64
	IsVegan bool
}

func main() {
    // Stage 1: Declare the 'openDays' array of 2 strings: "Saturday" and "Sunday"
	openDays := [2]string{"Saturday", "Sunday"}

    // Stage 2: Create the 'pantry' map (Flour: 10, Sugar: 5, Yeast: 0)
	pantry := map[string]int{"Flour": 10, "Sugar": 5, "Yeast": 0}

    // Stage 3: Use make() to create the empty 'shoppingList' string slice (len 0, cap 5)
	shoppingList := make([]string, 0, 5)

    // Stage 3: Loop over 'pantry'. If quantity is 0, append the key to 'shoppingList'
	for item, count := range pantry {
		if count == 0 {
			shoppingList = append(shoppingList, item)
		}
	}

    // Stage 4: Create the 'menu' slice of Recipes (one vegan, one not)
	menu := []Recipe{ {Name: "Chocolate Cake", Cost: 12.50, IsVegan: false}, {Name: "Vegan Tart", Cost: 15.00, IsVegan: true} }

    // Stage 4: Loop over 'menu'. If IsVegan is true, add Cost to 'veganCost'
	veganCost := 0.00
	for _, item := range menu {
		if item.IsVegan {
			veganCost+=item.Cost
		}
	}

    // Stage 5: Print openDays, shoppingList, and veganCost (using fmt.Printf for 2 decimals)
	fmt.Println("Open Days: ", openDays)
	fmt.Println("Shopping List: ", shoppingList)
	fmt.Printf("Vegan Production Cost: %.2f", veganCost)
}