// To declare a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
// the variable's type is inferred from the value on the right hand side.

// When the right hand side of the declaration is typed, the new variable is of that same type:

// for eg:1
// i := 42           // int
// f := 3.14         // float64
// g := 0.867 + 0.5i // complex128

// for eg:2
// var i int
// j := i // j is also an int

package main

import "fmt"

func main() {
	penniesPerText := 2.0

	// don't edit below this line
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}
