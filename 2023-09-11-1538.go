package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Exercise 3: Functions
	sum := add(4, 5)
	fmt.Printf("The sum of 4 and 5 is %d.\n", sum)

	// Exercise 4: Loops
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Exercise 5: Conditionals
	num := 10
	if num > 0 {
		fmt.Println("The number is positive.")
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	// Exercise 6: Arrays and Slices
	var languages [3]string
	languages[0] = "Go"
	languages[1] = "Python"
	languages[2] = "JavaScript"
	fmt.Println(languages)

	numbersSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(numbersSlice[2:4])

	// Exercise 7: Maps
	capitals := map[string]string{"USA": "Washington D.C.", "UK": "London", "India": "New Delhi"}
	fmt.Println(capitals["USA"])

	// Exercise 8: Structs
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person's name is %s and age is %d.\n", person.Name, person.Age)
}

func add(a, b int) int {
	return a + b
}