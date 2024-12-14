// Constants are declared with the const keyword. They can't use the := short declaration syntax.
//Value of constants cannot be updated in the code once it has been declared.

//for re1:
// const pi = 3.14159

// PROGRAM-1
// package main

// import "fmt"

// func main() {
// 	const premiumPlanName = "Premium Plan"
// 	const basicPlanName = "Basic Plan"

// 	// don't edit below this line

// 	fmt.Println("plan:", premiumPlanName)
// 	fmt.Println("plan:", basicPlanName)
// }

// PROGRAM-2
package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}




//declare multiple constants together
const (
	pi = 3.14159
	gravity = 9.81
)
