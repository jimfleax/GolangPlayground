/*
	Write a Go program named findian.go that performs the following steps:
	Prompt the user to enter a string.
	Check whether the string:
	The program should not be case-sensitive.
	If all conditions are met, print:  Found!
	Otherwise print: Not Found!

	Example Inputs
	These should print Found!
	ian
	Ian
	iuiygaygn
	I d skd a efju N
	These should print Not Found!
	ihhhhhn
	ina
	xian
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please, enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str := strings.ToLower(scanner.Text())

		condPre := strings.HasPrefix(str, "i")
		condSuf := strings.HasSuffix(str, "n")
		condCon := strings.Contains(str, "a")

		if condPre && condSuf && condCon {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}