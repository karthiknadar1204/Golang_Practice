package main

import "fmt"

func main() {
	// make(map[KeyType]ValueType)
	scores := make(map[string]int)

	// Adding key-value pairs
	scores["Alice"] = 98
	scores["Bob"] = 85
	fmt.Println(scores) // Output: map[Alice:98 Bob:85]
}
