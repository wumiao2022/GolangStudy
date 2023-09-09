package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var name string = "John Doe"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Example 3: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Example 4: Slices
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// Example 5: Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Functions
	add := func(a, b int) int {
		return a + b
	}
	result := add(3, 5)
	fmt.Println(result)

	// Example 7: Structs
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person)

	// Example 8: Concurrency
	ch := make(chan string)
	go func() {
		ch <- "Hello, Golang!"
	}()
	message := <-ch
	fmt.Println(message)
}