// package main

// import "fmt"

// func main() {
// 	day := "Monday"
// 	switch day {
// 	case "Monday":
// 		fmt.Println("Start of work week")
// 	case "Friday":
// 		fmt.Println("TGIF!")
// 	case "Saturday", "Sunday": // Multiple cases in one line
// 		fmt.Println("Weekend!")
// 	default:
// 		fmt.Println("Regular work day")
// 	}
// }



package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %v\n", v)
	case string:
		fmt.Printf("String: %v\n", v)
	case bool:
		fmt.Printf("Boolean: %v\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	checkType(42)      // Integer: 42
	checkType("Hello") // String: Hello
	checkType(true)    // Boolean: true
	checkType(3.14)    // Unknown type
}
