// 1. nil slices
// package main

// import "fmt"

// func main() {
// 	var slice []int //we could also write "num:[]int {}"
// 	fmt.Println(slice == nil) // Output: true
// 	fmt.Println(len(slice))   // Output: 0

// The cap() method returns the capacity of a slice, which is the maximum number of elements the slice can hold before it needs to be resized.
// 	fmt.Println(cap(slice))   // Output: 0

// 	// Appending to nil slice
// 	slice = append(slice, 1)

// 	fmt.Println(slice) // Output: [1]
// }







// 2. Using make()
// package main

// import "fmt"

// func main() {
// 	// make([]type, length, capacity)
// 	slice1 := make([]int, 3)    // length=3, capacity=3
// 	slice2 := make([]int, 3, 5) // length=3, capacity=5

// 	// When you append elements beyond current capacity, Go creates a new array with doubled capacity
// 	slice1 = append(slice1, 4, 6, 8, 10, 12, 14)
// 	fmt.Printf("slice1: %v, length: %d, capacity: %d\n", slice1, len(slice1), cap(slice1))
// 	fmt.Printf("slice2: %v, length: %d, capacity: %d\n", slice2, len(slice2), cap(slice2))
// 	fmt.Println("updated slice1 is ", slice1) //updated slice1 is  [0 0 0 4 6]
// }

// package main

// import "fmt"

// func main() {
// 	// Creating a slice with initial capacity
// 	numbers := make([]int, 0, 5)
// 	fmt.Printf("Initial - Length: %d, Capacity: %d\n", len(numbers), cap(numbers))
// 	fmt.Println("initial slice before looping is ", numbers)

// 	// Adding elements up to capacity
// 	for i := 1; i <= 5; i++ {
// 		numbers = append(numbers, i)
// 		fmt.Println("Slice during the Looping is ", numbers)
// 		fmt.Printf("After append %d - Length: %d, Capacity: %d\n",
// 			i, len(numbers), cap(numbers))
// 	}

// 	// Adding beyond capacity
// 	numbers = append(numbers, 6)
// 	fmt.Println("Slice after exceeding capacity", numbers)
// 	fmt.Printf("After exceeding capacity - Length: %d, Capacity: %d\n",
// 		len(numbers), cap(numbers))
// 	// Capacity doubles when exceeded
// }

// Output for the above:
// Initial - Length: 0, Capacity: 5
// initial slice before looping is  []
// Slice during the Looping is  [1]
// After append 1 - Length: 1, Capacity: 5
// Slice during the Looping is  [1 2]
// After append 2 - Length: 2, Capacity: 5
// Slice during the Looping is  [1 2 3]
// After append 3 - Length: 3, Capacity: 5
// Slice during the Looping is  [1 2 3 4]
// After append 4 - Length: 4, Capacity: 5
// Slice during the Looping is  [1 2 3 4 5]
// After append 5 - Length: 5, Capacity: 5
// Slice after exceeding capacity [1 2 3 4 5 6]
// After exceeding capacity - Length: 6, Capacity: 10










// copy method in slices
// package main
// import "fmt"
// func main() {
// 	src := []int{1, 2, 3}
// 	dst := make([]int, len(src))
// 	// Copy elements from src to dst
// 	copied := copy(dst, src)
// 	fmt.Printf("src: %v\n", src) // Output: src: [1 2 3]
// 	fmt.Printf("dst: %v\n", dst) // Output: dst: [1 2 3]
// 	fmt.Printf("copied elements: %d\n", copied) // Output: copied elements: 3
// }






//Slicing literal
// package main

// import "fmt"

// func main() {
//     // Basic array to demonstrate slicing
//     numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//     fmt.Println("Original:", numbers)

//     // Basic slicing examples
//     fmt.Println("numbers[1:5] =", numbers[1:5])    // [1 2 3 4]-->prints from [start index,end index-1]
//     fmt.Println("numbers[:3] =", numbers[:3])      // [0 1 2]-->prints till [end index-1] from the 0th index.
//     fmt.Println("numbers[4:] =", numbers[4:])      // [4 5 6 7 8 9]-->prints till last index starting from start index.
//     fmt.Println("numbers[:] =", numbers[:])        // [0 1 2 3 4 5 6 7 8 9]-->prints from start to end.

//     // More complex examples
//     slice1 := numbers[2:7]     // [2 3 4 5 6]
//     fmt.Println("\nslice1:", slice1)

//     // Slicing a slice
//     slice2 := slice1[1:3]     // [3 4]
//     fmt.Println("slice2:", slice2)

//     // Modifying elements
//     slice2[0] = 30            // This will affect all slices
//     fmt.Println("\nAfter modification:")
//     fmt.Println("Original numbers:", numbers)   // [0 1 2 30 4 5 6 7 8 9]
//     fmt.Println("slice1:", slice1)             // [2 30 4 5 6]
//     fmt.Println("slice2:", slice2)             // [30 4]
// }






//slices.equal method


//1. Basic Usage of slices.Equal
// package main

// import (
//     "fmt"
//     "slices" // Requires Go 1.21+
// )

// func main() {
//     // Comparing integer slices
//     slice1 := []int{1, 2, 3}
//     slice2 := []int{1, 2, 3}
//     slice3 := []int{1, 2, 4}

//     fmt.Println(slices.Equal(slice1, slice2))  // true
//     fmt.Println(slices.Equal(slice1, slice3))  // false

//     // Comparing string slices
//     strSlice1 := []string{"a", "b", "c"}
//     strSlice2 := []string{"a", "b", "c"}
//     strSlice3 := []string{"a", "b", "d"}

//     fmt.Println(slices.Equal(strSlice1, strSlice2))  // true
//     fmt.Println(slices.Equal(strSlice1, strSlice3))  // false
// }
// It compares elements in order
// Returns true if both slices have the same length and elements






// 2. Comparing Slices with Different Lengths
// package main

// import (
//     "fmt"
//     "slices"
// )

// func main() {
//     slice1 := []int{1, 2, 3}
//     slice2 := []int{1, 2, 3, 4}
//     slice3 := []int{1, 2}

//     fmt.Println(slices.Equal(slice1, slice2))  // false
//     fmt.Println(slices.Equal(slice1, slice3))  // false

//     // Empty slices
//     empty1 := []int{}
//     empty2 := []int{}
//     fmt.Println(slices.Equal(empty1, empty2))  // true
// }



// 3. Comparing Nil Slices
// package main

// import (
//     "fmt"
//     "slices"
// )

// func main() {
//     var slice1 []int        // nil slice
//     var slice2 []int        // nil slice
//     slice3 := []int{}       // empty slice, but not nil

//     fmt.Println(slices.Equal(slice1, slice2))  // true
//     fmt.Println(slices.Equal(slice1, slice3))  // true
    
//     // nil slices are considered equal to empty slices
//     fmt.Printf("slice1 is nil: %v\n", slice1 == nil)  // true
//     fmt.Printf("slice3 is nil: %v\n", slice3 == nil)  // false
// }






// 4. Comparing Custom Types
// package main

// import (
//     "fmt"
//     "slices"
// )

// type Person struct {
//     Name string
//     Age  int
// }

// func main() {
//     people1 := []Person{
//         {Name: "Alice", Age: 25},
//         {Name: "Bob", Age: 30},
//     }
    
//     people2 := []Person{
//         {Name: "Alice", Age: 25},
//         {Name: "Bob", Age: 30},
//     }

//     people3 := []Person{
//         {Name: "Alice", Age: 25},
//         {Name: "Bob", Age: 31},  // Different age
//     }

//     fmt.Println(slices.Equal(people1, people2))  // true
//     fmt.Println(slices.Equal(people1, people3))  // false
// }







// 5. Using with Different Types
package main

import (
    "fmt"
    "slices"
)

func main() {
    // Integer slices
    ints1 := []int{1, 2, 3}
    ints2 := []int{1, 2, 3}
    fmt.Println("Integers:", slices.Equal(ints1, ints2))

    // Float slices
    floats1 := []float64{1.1, 2.2, 3.3}
    floats2 := []float64{1.1, 2.2, 3.3}
    fmt.Println("Floats:", slices.Equal(floats1, floats2))

    // Byte slices
    bytes1 := []byte("hello")
    bytes2 := []byte("hello")
    fmt.Println("Bytes:", slices.Equal(bytes1, bytes2))
}