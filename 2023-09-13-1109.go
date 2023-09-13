package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var x int = 5
	var y float64 = 7.8
	var name string = "Golang"
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("name =", name)

	// Example 3: If statement
	num := 10
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// Example 4: Loop - for
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 5: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Example 6: Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Example 7: Functions
	result := add(3, 5)
	fmt.Println("Result =", result)
}

func add(a, b int) int {
	return a + b
}