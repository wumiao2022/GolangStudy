package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello World" to the console
	fmt.Println("Hello World")

	// Exercise 2: Create a variable called "message" and assign it the value "Hello"
	message := "Hello"
	fmt.Println(message)

	// Exercise 3: Create a slice of integers containing the values 1, 2, 3
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// Exercise 4: Create a function called "add" that takes two integer parameters and returns their sum
	fmt.Println(add(2, 3))

	// Exercise 5: Create a struct called "Person" with "name" and "age" fields
	type Person struct {
		name string
		age  int
	}
	person := Person{"John", 30}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}