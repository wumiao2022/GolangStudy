package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Declare and initialize a variable with value 10, then print its value
	num := 10
	fmt.Println(num)

	// Exercise 3: Create an array of integers with values 1, 2, 3, 4, 5 and print its elements
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// Exercise 4: Create a slice of integers with values 1, 2, 3, 4, 5 and print its elements
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Exercise 5: Create a map with key-value pairs of type string and int, then print the values using a loop
	myMap := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for _, value := range myMap {
		fmt.Println(value)
	}
}
