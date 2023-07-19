package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare and print a variable
	var name string = "John"
	fmt.Println("My name is", name)

	// Exercise 3: Perform addition and print the result
	var a, b int = 5, 3
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// Exercise 4: Use a loop to print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Calculate the factorial of a number
	factorial := 1
	num := 5
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("The factorial of", num, "is", factorial)
}