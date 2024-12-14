// You can declare multiple variables on the same line:

// mileage, company := 80276, "Tesla"

// The above is the same as:
// mileage := 80276
// company := "Tesla"

package main

import "fmt"

func main() {
	// declare here
	averageOpenRate, displayMessage := 50, "hi there"

	fmt.Println(averageOpenRate, displayMessage)
}
