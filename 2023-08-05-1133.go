package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Create a variable named "name" with value "John" and print it
	name := "John"
	fmt.Println(name)

	// Exercise 3: Create a function named add that takes two integers as parameters and returns their sum
	add := func(a, b int) int {
		return a + b
	}

	// Exercise 4: Call the add function with the values 5 and 3, and print the result
	result := add(5, 3)
	fmt.Println(result)

	// Exercise 5: Create a struct named "Person" with fields "name" and "age"
	type Person struct {
		name string
		age  int
	}

	// Exercise 6: Create a variable of type Person and initialize it with the values "Alice" and 25
	person := Person{name: "Alice", age: 25}

	// Exercise 7: Print the name and age of the person
	fmt.Println(person.name, person.age)
}