package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var age int = 25
	fmt.Printf("I am %d years old.\n", age)

	// Example 3: If-else statements
	score := 80

	if score >= 90 {
		fmt.Println("Excellent!")
	} else if score >= 70 {
		fmt.Println("Good!")
	} else {
		fmt.Println("Keep trying!")
	}

	// Example 4: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("The first number is %d.\n", numbers[0])

	// Example 5: Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Functions
	result := add(3, 4)
	fmt.Printf("The result of addition is %d.\n", result)

	// Example 7: Structs
	type Person struct {
		name    string
		age     int
		country string
	}

	person := Person{
		name:    "John",
		age:     30,
		country: "USA",
	}

	fmt.Println(person.name)

	// Example 8: Pointers
	x := 5
	y := &x
	fmt.Printf("The value of x is %d.\n", *y)
}

func add(a, b int) int {
	return a + b
}