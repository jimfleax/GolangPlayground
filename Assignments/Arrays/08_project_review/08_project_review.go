package main

import "fmt"

// Stage 1: Define the Bike struct here
type Bike struct {
	ID       string
	Rate     float64
	IsRented bool
}

func main() {
	// Stage 1: Declare the 'branches' array of 2 strings: "Downtown" and "Campus"
	branches := [2]string{"Downtown", "Campus"}

	// Stage 2: Create the 'fleet' slice of Bikes (one rented, one not)
	fleet := []Bike{{ID: "B1", Rate: 15.00, IsRented: true}, {ID: "B2", Rate: 20.00, IsRented: false}}

	// Stage 2: Create the 'rentals' map ("Alice" -> "B1")
	rentals := map[string]string{"Alice": "B1"}

	// Stage 3: Use make() to create the 'revenueLog' slice of float64 (len 0, cap 5)
	revenueLog := make([]float64, 0, 5)

	// Stage 4: Loop over 'fleet'. If IsRented is true, append Rate to 'revenueLog'
	for _, bike := range fleet {
		if bike.IsRented {
			revenueLog = append(revenueLog, bike.Rate)
		}
	}

	// Stage 5: Loop over 'revenueLog' to sum totalRevenue. Print branches, rentals, and totalRevenue.
	totalRevenue := 0.0
	for _, rate := range revenueLog {
		totalRevenue += rate
	}

	fmt.Println("Branches: ", branches)
	fmt.Println("Rentals: ", rentals)
	fmt.Println("Total Revenue: ", totalRevenue)
}
