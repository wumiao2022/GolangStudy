package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare and assign a string variable
	message := "I love Golang"

	// Exercise 3: Print the value of the string variable
	fmt.Println(message)

	// Exercise 4: Declare and assign an integer variable
	number := 10

	// Exercise 5: Print the value of the integer variable
	fmt.Println(number)

	// Exercise 6: Declare a constant variable
	const pi = 3.14159

	// Exercise 7: Print the value of the constant variable
	fmt.Println(pi)

	// Exercise 8: Create a slice with 3 elements
	numbers := []int{1, 2, 3}

	// Exercise 9: Print the elements of the slice
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 10: Create a map with string keys and integer values
	person := map[string]int{
		"John":  25,
		"Jane":  30,
		"Alice": 28,
	}

	// Exercise 11: Print the values of the map
	for name, age := range person {
		fmt.Println(name, age)
	}
}
