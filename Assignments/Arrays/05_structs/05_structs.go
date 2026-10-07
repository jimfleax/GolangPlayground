package main

import "fmt"

// Stage 1: Define the RepairJob struct here
type RepairJob struct {
	Customer string
	Gadget string
	Cost float64
	IsFixed bool
}

func main() {
    // Stage 2: Initialize a RepairJob for John Doe's Laptop ($150) and print it
	job := RepairJob{"John Doe", "Laptop", 150.00, false}
	fmt.Println("Initial Ticket:", job)
    // Stage 3: Add 45.50 to the Cost, mark IsFixed as true, and print the final ticket and bill
	job.Cost+=45.50
	job.IsFixed = true
	fmt.Println("Final Ticket", job)
	fmt.Printf("Bill: %s owes %.2f for the %s repair.", job.Customer, job.Cost, job.Gadget)
}