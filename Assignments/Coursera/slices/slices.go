/*
	Write a Go program named slice.go that performs the following steps:
	Create an empty integer slice with length 3 before entering the loop.
	Prompt the user to enter integers one at a time.
	Each time an integer is entered:
	The slice must expand dynamically to hold additional integers.
	The program should continue running until the user enters “X” instead of a number.

	Example behavior:
	Input Sequence
	3
	1
	5
	Output after each entry
	[3]
	[1 3]
	[1 3 5]
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := make([]int, 0,3)
	var input string
	for {
		fmt.Print("Enter a number: ")
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		input, _ := strconv.Atoi(input)
		nums = append(nums, input)
		sort.Ints(nums)
		fmt.Println(nums)
	}
}
