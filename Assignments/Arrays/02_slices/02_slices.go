package main

import "fmt"

func main() {
    // 1. Declare a slice of strings named 'fruits' using a slice literal ("Apple", "Banana")
	fruits := []string{"Apple", "Banana"}
    // 2. Append "Orange" to the fruits slice
	fruits = append(fruits, "Orange")

    // 3. Print the slice and its length using fmt.Println and len()
	fmt.Println(fruits, len(fruits))

}