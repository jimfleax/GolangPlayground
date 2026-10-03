package main

import (
	"fmt"
	"strings"
)

func main() {
	csvData := "apple,banana,cherry"
	
	// 1. Split csvData by comma into fruitSlice
	fruitSlice := strings.Split(csvData, ",")
	// 2. Print fruitSlice
	fmt.Println(fruitSlice)
	// 3. Join fruitSlice using "- " into hyphenatedFruits
	hyphenatedFruits := strings.Join(fruitSlice, "-")
	// 4. Print hyphenatedFruits
	fmt.Println(hyphenatedFruits)
}