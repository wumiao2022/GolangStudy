package main

import (
	"fmt"
)

func main() {
	// Example 1: Printing "Hello, World!"
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var x int = 5
	const y float64 = 3.14

	// Example 3: Control Structures - If-Else
	if x < 10 {
		fmt.Println("x is less than 10")
	} else {
		fmt.Println("x is greater than or equal to 10")
	}

	// Example 4: Control Structures - For loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Arrays and Slices
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	slice := []int{4, 5, 6}

	fmt.Println(arr)
	fmt.Println(slice)

	// Example 6: Functions
	result := add(2, 3)
	fmt.Println(result)

	// Example 7: Pointers
	var ptr *int
	value := 10
	ptr = &value
	fmt.Println(*ptr)

	// Example 8: Structs
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "John", Age: 25}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}