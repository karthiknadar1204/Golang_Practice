// Variables are declared using the var keyword. For
// example, to declare a variable called number of type int, you would write:

// var number int

package main

import "fmt"

func main() {
	var smsSendingLimit int = 50
	var costPerSMS float64 = 0.10
	var hasPermission bool = true
	var username string = "Syl"

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
