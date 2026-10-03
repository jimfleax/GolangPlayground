package main

import (
	"fmt"
	"strings"
)

func main() {
	systemMessage := "Error: Connection timeout"
	
	// 1. Convert to uppercase and print
	fmt.Println(strings.ToUpper(systemMessage))
	
	// 2. Check if it contains "Error" and print the boolean result
	fmt.Println(strings.Contains(systemMessage, "Error"))
	
}