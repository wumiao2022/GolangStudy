package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variable declaration and initialization
	var num1 int = 10
	var num2 float64 = 3.14
	var str string = "Golang"
	fmt.Println(num1, num2, str)

	// Example 3: Simple arithmetic operation
	result := num1 + int(num2)
	fmt.Println("Result of addition:", result)

	// Example 4: Conditional statement
	if result > 15 {
		fmt.Println("Result is greater than 15")
	} else {
		fmt.Println("Result is not greater than 15")
	}

	// Example 5: Looping
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i+1)
	}

	// Example 6: Function
	fmt.Println("Sum of two numbers:", sum(5, 3))

	// Example 7: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Third element of the array:", numbers[2])

	// Example 8: Slices
	slice := numbers[1:4]
	fmt.Println("Slice:", slice)

	// Example 9: Maps
	ages := map[string]int{
		"John":  25,
		"Emily": 30,
		"Jack":  27,
	}
	fmt.Println("Age of John:", ages["John"])
}

// Function to calculate the sum of two numbers
func sum(a, b int) int {
	return a + b
}