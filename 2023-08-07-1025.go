package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare and print a string variable
	message := "Welcome to Golang!"
	fmt.Println(message)

	// Exercise 3: Perform arithmetic operations
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 4: Use if-else condition to check a number
	number := 7
	if number%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// Exercise 5: Create and print a slice of integers
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Exercise 6: Iterate over a slice and print each element
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Exercise 7: Create a function and call it
	greet("John")

	// Exercise 8: Create a struct and print its properties
	person := Person{
		Name:    "Alice",
		Age:     25,
		Country: "USA",
	}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Country:", person.Country)

	// Exercise 9: Define an interface and implement it
	circle := Circle{radius: 5.0}
	printArea(circle)
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

type Person struct {
	Name    string
	Age     int
	Country string
}

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(shape Shape) {
	fmt.Println("Area:", shape.Area())
}