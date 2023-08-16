package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var name string = "John"
	var age int = 28
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Exercise 3: Arithmetic Operations
	var a, b int = 10, 5
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	remainder := a % b
	fmt.Printf("Sum: %d, Difference: %d, Product: %d, Quotient: %d, Remainder: %d\n", sum, difference, product, quotient, remainder)

	// Exercise 4: Conditionals
	num := 10
	if num > 0 {
		fmt.Println("Number is positive")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// Exercise 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Arrays
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 7: Slices
	numbersSlice := numbers[1:4]
	fmt.Println(numbersSlice)

	// Exercise 8: Maps
	citiesPopulation := map[string]int{
		"New York":      8622698,
		"Los Angeles":   3990456,
		"San Francisco": 883305,
	}
	fmt.Println(citiesPopulation)
}