package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var num1 int = 10
	var num2 float64 = 3.14
	var str string = "Golang"
	fmt.Printf("num1: %d, num2: %.2f, str: %s\n", num1, num2, str)

	// Exercise 3: Constants
	const PI float64 = 3.1415
	const companyName string = "ABC Corp"
	fmt.Println("PI:", PI)
	fmt.Println("Company Name:", companyName)

	// Exercise 4: Arithmetic Operators
	a := 10
	b := 5
	add := a + b
	subtract := a - b
	multiply := a * b
	divide := a / b
	remainder := a % b
	fmt.Println("Addition:", add)
	fmt.Println("Subtraction:", subtract)
	fmt.Println("Multiplication:", multiply)
	fmt.Println("Division:", divide)
	fmt.Println("Remainder:", remainder)

	// Exercise 5: Conditional Statements
	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else {
		fmt.Println("num1 is equal to num2")
	}

	// Exercise 6: Loops
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Exercise 7: Arrays
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Numbers:", numbers)

	// Exercise 8: Slices
	numbersSlice := numbers[1:4]
	fmt.Println("Slice:", numbersSlice)

	// Exercise 9: Maps
	salary := make(map[string]int)
	salary["John"] = 5000
	salary["Jane"] = 6000
	salary["Mike"] = 7000
	fmt.Println("Salary:", salary)

	// Exercise 10: Functions
	result := multiplyNumbers(5, 6)
	fmt.Println("Result:", result)
}

func multiplyNumbers(x, y int) int {
	return x * y
}