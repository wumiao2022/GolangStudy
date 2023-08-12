package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare two variables, x and y, with values 10 and 5 respectively. Print their sum to the console.
	x := 10
	y := 5
	sum := x + y
	fmt.Println(sum)

	// Exercise 3: Declare a variable called name and assign it your name. Print "Hello, <name>!" to the console.
	name := "John"
	fmt.Println("Hello,", name, "!")

	// Exercise 4: Create a function called multiply that takes two integers as parameters and returns their product.
	// Call the function with some values and print the result.
	multiply(2, 3)
	multiply(5, 10)

	// Exercise 5: Given a list of numbers [1, 2, 3, 4, 5], iterate over the list and print each number to the console.
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 6: Create a function called isEven that takes an integer as a parameter and returns true if the number is even, otherwise false.
	// Call the function with some values and print the result.
	fmt.Println(isEven(2))
	fmt.Println(isEven(5))
}

func multiply(x, y int) {
	fmt.Println(x * y)
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
```