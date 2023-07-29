package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable "name" and assign your name to it, then print it to the console
	name := "John"
	fmt.Println(name)

	// Exercise 3: Create a function called "add" that takes in two integers and returns their sum.
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))

	// Exercise 4: Create a slice of integers and print its length and elements
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(len(numbers))
	fmt.Println(numbers)

	// Exercise 5: Create a map that maps string keys to integer values. Add some key-value pairs and print them.
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap)

	// Exercise 6: Use a for loop to print the numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 7: Use a while loop to print the even numbers from 2 to 20
	num := 2
	for num <= 20 {
		fmt.Println(num)
		num += 2
	}

	// Exercise 8: Write a recursive function to calculate the factorial of a number
	factorial := 1
	number := 5
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d is %d\n", number, factorial)
}