// Inside a function (like the main function) the := short assignment statement can be used in place of a
//  var declaration. The := operator infers the type of the new variable based on the value.
//   It's colloquially called the walrus operator because it looks like a walrus... sort of.

package main

import "fmt"

func main() {
	messageStart := "I am "
	age := 30
	messageEnd := " years old"
	fmt.Println(messageStart, age, messageEnd)
}
