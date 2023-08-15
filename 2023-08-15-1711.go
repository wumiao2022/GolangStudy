package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var a int = 10
	const b float64 = 3.14

	fmt.Println(a)
	fmt.Println(b)

	// Example 3: Arithmetic Operations
	fmt.Println(a + 5)
	fmt.Println(a - 3)
	fmt.Println(a * 2)
	fmt.Println(a / 3)

	// Example 4: Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 5: If-else statement
	if a < 5 {
		fmt.Println("a is less than 5")
	} else {
		fmt.Println("a is greater than or equal to 5")
	}

	// Example 6: Functions
	result := add(3, 5)
	fmt.Println(result)

	// Example 7: Arrays and Slices
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]

	fmt.Println(arr)
	fmt.Println(slice)

	// Example 8: Structs
	type Person struct {
		name   string
		age    int
		gender string
	}

	person := Person{name: "John", age: 25, gender: "Male"}

	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}