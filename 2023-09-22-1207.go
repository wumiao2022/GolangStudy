package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variable Declaration
	var message string = "Hello, Golang!"
	fmt.Println(message)

	// Exercise 3: Constant Declaration
	const pi = 3.14159
	fmt.Println(pi)

	// Exercise 4: Arithmetic Operations
	num1 := 10
	num2 := 5
	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)

	// Exercise 5: If-Else Statement
	num := 10
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercise 6: For Loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 7: Switch Statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Midweek")
	}

	// Exercise 8: Array Declaration
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)

	// Exercise 9: Slice Declaration
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// Exercise 10: Map Declaration
	capitals := map[string]string{
		"China":  "Beijing",
		"USA":    "Washington D.C.",
		"France": "Paris",
	}
	fmt.Println(capitals)
}