package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare and print variables
	var num1 int = 10
	num2 := 5
	fmt.Printf("num1: %d, num2: %d\n", num1, num2)

	// Exercise 3: Perform basic arithmetic operations
	addition := num1 + num2
	subtraction := num1 - num2
	multiplication := num1 * num2
	division := num1 / num2
	fmt.Printf("Addition: %d, Subtraction: %d, Multiplication: %d, Division: %d\n", addition, subtraction, multiplication, division)

	// Exercise 4: Perform increment and decrement operations
	num1++
	num2--
	fmt.Printf("Increment num1: %d, Decrement num2: %d\n", num1, num2)

	// Exercise 5: Use if-else statements
	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num2 is greater than num1")
	} else {
		fmt.Println("num1 is equal to num2")
	}

	// Exercise 6: Use switch statements
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	// Exercise 7: Use for and range loop
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 8: Create and call a function
	result := addNumbers(10, 20)
	fmt.Println("Sum:", result)
}

func addNumbers(a, b int) int {
	return a + b
}