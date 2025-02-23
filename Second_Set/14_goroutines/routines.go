// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printNumbers() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Millisecond * 500) // Simulate some work
// 	}
// }

// func main() {
// 	go printNumbers() // Starts a new goroutine
// 	fmt.Println("Hello from main!")

// 	time.Sleep(time.Second * 3) // Give goroutine time to finish
// 	fmt.Println("Main function finished")
// }







package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Millisecond * 300) // Simulate work
	}
}

func main() {
	go task("Goroutine 1")
	go task("Goroutine 2")
	go task("Goroutine 3")

	time.Sleep(time.Second * 2) // Wait for goroutines to finish
	fmt.Println("All tasks done")
}
