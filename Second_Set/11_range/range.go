package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	//iterating using loop without range.
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// iterating using loop with range
	sum := 0
	// for index,value := range <data_structure>{}
	//we can ignore the index by using _(underscore) in place of it.
	//index starts from 0.
	for i, num := range nums {
		sum = sum + num
		fmt.Println("nums here are", num)
		fmt.Println("index are", i)
	}
	fmt.Println(sum)

	//range in map
	m := map[string]string{"fname": "john", "lname": "doe"}
	for key, value := range m {
		fmt.Println("Key-value pairs in the map are:", key, value)
	}

	//range in strings
	for i, c := range "golang" {

		//prints the ascii value of the string.
		fmt.Println(i,c)

		//typecast value to string
		fmt.Println(i, string(c))
	}
	    // Method 1: Using string() conversion
		fmt.Println("Method 1: Using string() conversion")
		for i, c := range "golang" {
			fmt.Printf("index: %d, character: %s, rune value: %d\n", i, string(c), c)
		}
	
		// Method 2: Using fmt.Sprintf
		fmt.Println("\nMethod 2: Using fmt.Sprintf")
		for i, c := range "golang" {
			char := fmt.Sprintf("%c", c)
			fmt.Printf("index: %d, character: %s, rune value: %d\n", i, char, c)
		}
}
