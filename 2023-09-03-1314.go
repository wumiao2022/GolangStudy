package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variable declaration and initialization
	var name string = "John"
	var age int = 25
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// Exercise 3: Arithmetic operations
	var a = 10
	var b = 5
	var sum = a + b
	var difference = a - b
	var product = a * b
	var quotient = a / b
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Exercise 4: If-else statement
	var num = 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Exercise 5: For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Slice creation and iteration
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 7: Function declaration and call
	result := sumTwoNumbers(2, 3)
	fmt.Println("Sum:", result)
}

func sumTwoNumbers(a, b int) int {
	return a + b
}