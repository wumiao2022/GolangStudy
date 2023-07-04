package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Create a variable and assign it with a string value, then print its value
	var name string = "Alice"
	fmt.Println("My name is", name)

	// Exercise 3: Create an array of integers, initialize it with some values, then print each element
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println(number)
	}

	// Exercise 4: Create a function that takes two integers as parameters and returns their sum
	fmt.Println("Sum of 3 and 5 is", sum(3, 5))

	// Exercise 5: Create a struct type for a person, with name and age fields, then create an object of that type and print its fields
	type Person struct {
		name string
		age  int
	}

	person := Person{name: "Bob", age: 30}
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}

func sum(a, b int) int {
	return a + b
}
