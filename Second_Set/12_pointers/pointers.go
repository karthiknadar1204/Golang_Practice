package main

import "fmt"

//by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("in changeNum", num)

// }

// by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("in changeNum", *num)
}

func main() {
	num := 1
	// changeNum(num)--->by value
	fmt.Println("Memory Address", &num)
	changeNum(&num) //---->by reference

	fmt.Println("After change in main", num)
}
