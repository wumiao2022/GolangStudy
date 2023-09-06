package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable "name" and assign it with your name. Then print the value of the variable.
	name := "John Doe"
	fmt.Println(name)

	// Exercise 3: Create a function called "add" that takes two integers as parameters and returns their sum.
	fmt.Println(add(3, 5))

	// Exercise 4: Create a "for" loop that prints the numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Create a slice of integers with values [1, 2, 3, 4, 5]. Then print the sum of all the numbers.
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}