// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Mark this goroutine as done when finished
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // Simulating work
// 	fmt.Printf("Worker %d done\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup // Create a WaitGroup

// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1) // Add 1 to the WaitGroup for each goroutine
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() // Wait for all goroutines to complete
// 	fmt.Println("All workers finished!")
// }










package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		wg.Add(1) // Increment the wait count

		go func(n string) {
			defer wg.Done()
			fmt.Println("Hello,", n)
		}(name) // Pass name as parameter to avoid closure issues
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All greetings sent!")
}
