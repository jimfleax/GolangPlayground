package main

import "fmt"

func main() {
	userFirstName := "Reetabrata"
	userLastName := "Bhandari"
	completeDisplayName := userFirstName + " " + userLastName
	length := len(completeDisplayName)
	fmt.Println(completeDisplayName, length)
}