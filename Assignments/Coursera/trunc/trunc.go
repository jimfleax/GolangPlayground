/*
	Write a Go program named trunc.go that performs the following steps:
	Prompt the user to enter a floating point number.
	Read the number entered by the user.
	Print the truncated integer value of the number.

	Example behavior:
	Input
	3.7
	Output
	3
	Input
	-2.9
	Output
	-2
*/

package main

import (
	"fmt"
)

func main() {
	var input float64
	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&input)
	fmt.Printf("The truncated integer value is %d", int(input))
}
