package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare a variable of type int and print its value
	var num int = 42
	fmt.Println(num)

	// Exercise 3: Create a function that adds two numbers and returns their sum
	sum := add(10, 20)
	fmt.Println(sum)

	// Exercise 4: Create a slice of strings and print each element on a new line
	strings := []string{"apple", "banana", "cherry"}

	for _, str := range strings {
		fmt.Println(str)
	}

	// Exercise 5: Create a map to store names and ages, and print each key-value pair
	ages := map[string]int{
		"John": 25,
		"Jane": 30,
		"Bob":  35,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

func add(a, b int) int {
	return a + b
}
