package main

import "fmt"

func main() {
    // 1. Declare an array of integers named 'scores' with length 3
	var scores [3]int;

    // 2. Assign values 85, 92, and 78 to the respective indices
	scores[0] = 85
	scores[1] = 92
	scores[2] = 78

    // 3. Print the array and its length using fmt.Println and len()
	fmt.Println(scores, len(scores))

}