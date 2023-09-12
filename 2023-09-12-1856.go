package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console.
	fmt.Println("Hello, World!")

	// Exercise 2: Define and print a variable with the value 10.
	number := 10
	fmt.Println(number)

	// Exercise 3: Create a function that takes two integers as parameters and returns their sum.
	sum := addNumbers(5, 3)
	fmt.Println(sum)

	// Exercise 4: Create a function that takes a string as a parameter and prints it in uppercase.
	printUppercase("hello")

	// Exercise 5: Create a struct type called "Person" with "name" and "age" fields. Create an instance of this struct and print its values.
	person := Person{name: "John", age: 25}
	fmt.Println(person.name, person.age)
}

func addNumbers(a, b int) int {
	return a + b
}

func printUppercase(str string) {
	fmt.Println(strUpper(str))
}

func strUpper(str string) string {
	return strings.ToUpper(str)
}

type Person struct {
	name string
	age  int
}