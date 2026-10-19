package main

import "fmt"

func main() {
    // Stage 1: Create the 'applicants' slice of strings
	applicants := []string{"Alice", "Bob", "Charlie"}

    // Stage 2: Loop over 'applicants' and print the plot assignments
	for index, name := range applicants {
		fmt.Printf("Plot %d assigned to %s\n", index+1, name)
	}

    // Stage 3: Create the 'waterUsage' map and add the two plots
	waterUsage := map[string]float64{"Plot A": 12.5, "Plot B": 8.0}

    // Stage 4: Loop over 'waterUsage' and print the liters used per plot
	for plot, litres := range waterUsage {
		fmt.Printf("%s used %.1f litres of water.\n", plot, litres)
	}
}