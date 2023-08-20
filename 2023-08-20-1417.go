package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var x int = 10
	const y float64 = 3.14

	// Example 3: Functions
	func add(a, b int) int {
		return a + b
	}

	sum := add(x, 5)
	fmt.Println("Sum:", sum)

	// Example 4: Arrays and Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	slice := numbers[1:3]
	fmt.Println("Slice:", slice)

	// Example 5: Conditional Statements
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// Example 6: Loops
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Number:", numbers[i])
	}

	// Example 7: Maps
	ages := map[string]int{
		"John":   25,
		"Emma":   30,
		"Robert": 40,
	}
	fmt.Println("Ages:", ages)

	// Example 8: Pointers
	var ptr *int
	ptr = &x
	fmt.Println("Value of x:", *ptr)

	// Example 9: Structs
	type Person struct {
		name string
		age  int
	}

	person := Person{"Alice", 20}
	fmt.Println("Person:", person)

	// Example 10: Goroutines
	go func() {
		fmt.Println("Hello from goroutine!")
	}()

	fmt.Scanln()
}