package main

import "fmt"

func main() {
    // 1. Use make() to create an integer slice named 'numbers' with length 5, and print it
	numbers := make([]int, 5)
	fmt.Println(numbers)

    // 2. Create a slice of strings named 'colors' with "Red", "Green", "Blue", "Yellow"
	colors := []string{"Red", "Green", "Blue", "Yellow"}

    // 3. Slice 'colors' to extract the first three elements into 'primaryColors', and print it
	primaryColors := colors[:3]
	fmt.Println(primaryColors)
}