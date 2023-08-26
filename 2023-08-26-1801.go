package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, world!"
	fmt.Println("Hello, world!")

	// Exercise 2: Add two numbers and print the result
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Create a slice, append elements, and print the slice
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8)
	fmt.Println("Slice:", numbers)

	// Exercise 4: Iterate over a string and print each character
	word := "Golang"
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}

	// Exercise 5: Create a map, add key-value pairs, and print the map
	person := map[string]string{
		"name":  "John Doe",
		"email": "johndoe@example.com",
		"age":   "30",
	}
	fmt.Println("Person:", person)
}