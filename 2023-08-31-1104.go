package main

import "fmt"

func main() {
	// Exercise 1: Declare two variables, "a" and "b", and assign values to them. Print the sum of "a" and "b".
	a := 10
	b := 5
	fmt.Println(a + b)

	// Exercise 2: Declare a string variable named "message" and assign a greeting message to it. Print the message.
	message := "Hello, world!"
	fmt.Println(message)

	// Exercise 3: Declare a slice containing the numbers 1 to 5. Print all the numbers in the slice.
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 4: Declare a map called "student" with keys "name" and "age". Assign values to the keys and print them.
	student := map[string]interface{}{
		"name": "John Doe",
		"age":  20,
	}
	fmt.Println(student["name"])
	fmt.Println(student["age"])

	// Exercise 5: Write a function called "multiply" that takes two integers as parameters and returns their product.
	fmt.Println(multiply(3, 4))
}

func multiply(a, b int) int {
	return a * b
}