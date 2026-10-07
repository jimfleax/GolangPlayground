package main

import "fmt"

func main() {
    // 1. Create a map named 'inventory' with string keys and int values (Apples: 10, Oranges: 5)
	inventory := map[string]int{"Apples": 10, "Organes": 5}
    // 2. Add "Bananas" with a value of 20 to the map
	inventory["Bananas"] = 20
    // 3. Delete "Oranges" from the map using delete(), then print the map using fmt.Println
	delete(inventory, "Oranges")
	fmt.Println(inventory)
}