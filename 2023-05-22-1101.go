package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!" to console
	fmt.Println("Hello, World!")

	// Example 2: Declare and print a variable
	name := "John"
	fmt.Printf("My name is %s\n", name)

	// Example 3: Create a function that returns the sum of two numbers
	sum := add(3, 5)
	fmt.Printf("The sum of 3 and 5 is %d\n", sum)

	// Example 4: Create a struct and print its values
	person := Person{name: "Alice", age: 25, city: "New York"}
	fmt.Printf("Name: %s\nAge: %d\nCity: %s\n", person.name, person.age, person.city)

	// Example 5: Create a map and print its values
	scores := map[string]int{"John": 90, "Alice": 95, "Bob": 80}
	fmt.Printf("John's score is %d\n", scores["John"])
}

type Person struct {
	name string
	age  int
	city string
}

func add(a, b int) int {
	return a + b
}