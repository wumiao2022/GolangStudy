package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello World" on the console
	fmt.Println("Hello World")

	// Exercise 2: Declare and initialize a variable "name" with your name, and print it on the console
	name := "John Doe"
	fmt.Println(name)

	// Exercise 3: Calculate the square of a given number and print the result on the console
	number := 7
	square := number * number
	fmt.Println(square)

	// Exercise 4: Create a slice of integers and print its length and capacity on the console
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// Exercise 5: Create a map with key-value pairs of book titles and their authors, and print the map on the console
	books := map[string]string{
		"Book 1": "Author 1",
		"Book 2": "Author 2",
		"Book 3": "Author 3",
	}
	fmt.Println(books)
}