package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Print the numbers 1 to 10 using a loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 3: Find the sum of all even numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of even numbers:", sum)

	// Exercise 4: Convert a string to uppercase
	stringToConvert := "hello, world!"
	upperCaseString := ""
	for _, char := range stringToConvert {
		upperChar := string(char - 32)
		upperCaseString += upperChar
	}
	fmt.Println("Uppercase string:", upperCaseString)

	// Exercise 5: Calculate the factorial of a number
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, ":", factorial)
}