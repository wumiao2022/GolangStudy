package main

import "fmt"

func main() {
	// Exercise 1
	fmt.Println("Hello, World!")

	// Exercise 2
	name := "John"
	age := 25
	fmt.Println("My name is", name, "and I am", age, "years old.")

	// Exercise 3
	sum := add(5, 10)
	fmt.Println("The sum of 5 and 10 is", sum)

	// Exercise 4
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("The length of the array is", len(arr))

	// Exercise 5
	sayHello("Alice")
}

func add(a, b int) int {
	return a + b
}

func sayHello(name string) {
	fmt.Println("Hello,", name)
}