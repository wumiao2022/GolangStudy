package main

import "fmt"

func main() {
	// Exercise 1: Hello, World!
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var num1 int
	var num2 float64
	var str string

	num1 = 10
	num2 = 3.14
	str = "Golang"

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(str)

	// Exercise 3: Operators
	a := 5
	b := 7

	sum := a + b
	subtraction := a - b
	multiplication := a * b
	division := a / b
	remainder := a % b

	fmt.Println(sum)
	fmt.Println(subtraction)
	fmt.Println(multiplication)
	fmt.Println(division)
	fmt.Println(remainder)

	// Exercise 4: Functions
	result := add(3, 4)
	fmt.Println(result)

	// Exercise 5: Conditional Statements
	score := 85

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	// Exercise 6: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 7: Arrays
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println(numbers)

	// Exercise 8: Slices
	names := []string{"Alice", "Bob", "Charlie"}

	fmt.Println(names)

	// Exercise 9: Maps
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	fmt.Println(ages)

	// Exercise 10: Pointers
	x := 10
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
}

func add(a, b int) int {
	return a + b
}