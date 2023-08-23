package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable "name" and assign it your name, then print it.
	var name string = "John"
	fmt.Println(name)

	// Exercise 3: Create an array "numbers" with elements 1, 2, 3, 4, 5 and print it.
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 4: Create a function "add" that takes in two integer parameters and returns their sum.
	fmt.Println(add(10, 20))

	// Exercise 5: Create a struct "Person" with fields "name" and "age", initialize it with values and print it.
	person := Person{name: "Alice", age: 25}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}