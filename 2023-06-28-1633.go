package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variable Declaration and Initialization
	var num1 int = 10
	num2 := 5

	// Exercise 3: Basic Arithmetic Operations
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	// Exercise 4: Printing the Results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Exercise 5: Control Flow - if-else
	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else {
		fmt.Println("num1 is equal to num2")
	}

	// Exercise 6: Control Flow - for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Exercise 7: Control Flow - switch
	day := "Wednesday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")
	case "Friday":
		fmt.Println("Today is Friday")
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Exercise 8: Arrays
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Exercise 9: Slices
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)

	for _, value := range slice {
		fmt.Println(value)
	}

	// Exercise 10: Maps
	ages := map[string]int{
		"John":  25,
		"Mary":  30,
		"Peter": 35,
	}

	fmt.Println("Mary's age is", ages["Mary"])

	// Exercise 11: Functions
	result := multiply(3, 4)
	fmt.Println("Result:", result)
}

func multiply(a, b int) int {
	return a * b
}