package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console.
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable x with a value of 5 and print its value.
	x := 5
	fmt.Println(x)

	// Exercise 3: Declare a constant y with a value of 10 and print its value.
	const y = 10
	fmt.Println(y)

	// Exercise 4: Declare a function called "add" that takes two integer parameters, 
	// adds them together, and returns the result. Print the result to the console.
	result := add(3, 4)
	fmt.Println(result)

	// Exercise 5: Create a slice called "numbers" and add the values 1, 2, and 3 to it.
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// Exercise 6: Use a for loop to iterate through the slice "numbers" and print each value to the console.
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func add(a, b int) int {
	return a + b
}
