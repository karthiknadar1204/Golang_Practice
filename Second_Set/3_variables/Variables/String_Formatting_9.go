// Go follows the printf tradition from the C language

//DEFAULT REPRESENTATION->The %v variant prints the Go syntax representation of a value, it's a nice default.
// s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

// s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old

//STRING->%s
// s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old

//INTEGER->%d
// s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old

// FLOAT-> %.nf
// s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
// s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old

package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hii %s, your open rate is %.2f percent", name, openRate)

	// don't edit below this line

	fmt.Print(msg)
}
