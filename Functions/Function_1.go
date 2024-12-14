// To make code easier to read, the variable type comes after the variable name.

// func sub(x int, y int) int {
//   return x-y
// }

// x int
// p *int
// a [3]int

package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}






// Ignoring Return Values
// A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore,
// or more precisely, the blank identifier _.

// For example:

// func getPoint() (x int, y int) {
//   return 3, 4
// }

// ignore y value
// x, _ := getPoint()

// Even though getPoint() returns two values, we can capture the first one and ignore the second. In Go, the blank identifier isn't
// just a convention; it's a real language feature that completely discards the value.






// func getCoords() (x, y int){
// 	// x and y are initialized with zero values
  
// 	return // automatically returns x and y
//   }

//   // Is the same as:
  
//   func getCoords() (int, int){
// 	var x int
// 	var y int
// 	return x, y
//   }

// In the first example, x and y are the return values. At the end of the function, we could simply write return to return the values 
// of those two variables, rather than writing return x,y.

















// Early Returns
// Go supports the ability to return early from a function. This is a powerful feature that can clean up code, especially when used as guard clauses.

//Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. Instead of using if/else chains, we just return early from the function at the end of each conditional block.

// func divide(dividend, divisor int) (int, error) {
// 	if divisor == 0 {
// 		return 0, errors.New("Can't divide by zero")
// 	}
// 	return dividend/divisor, nil
// }

// Error handling in Go naturally encourages developers to make use of guard clauses. When I started writing more JavaScript, I
//I was disappointed to see how many nested conditionals existed in the code I was working on.

// Let’s take a look at an exaggerated example of nested conditional logic:

// func getInsuranceAmount(status insuranceStatus) int {
//   amount := 0
//   if !status.hasInsurance(){
//     amount = 1
//   } else {
//     if status.isTotaled(){
//       amount = 10000
//     } else {
//       if status.isDented(){
//         amount = 160
//         if status.isBigDent(){
//           amount = 270
//         }
//       } else {
//         amount = 0
//       }
//     }
//   }
//   return amount
// }

// This could be written with guard clauses instead:

// func getInsuranceAmount(status insuranceStatus) int {
//   if !status.hasInsurance(){
//     return 1
//   }
//   if status.isTotaled(){
//     return 10000
//   }
//   if !status.isDented(){
//     return 0
//   }
//   if status.isBigDent(){
//     return 270
//   }
//   return 160
// }