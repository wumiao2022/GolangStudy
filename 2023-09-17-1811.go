package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Printing variables
	name := "John"
	age := 30
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Example 3: Looping
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 4: Arrays and Slices
	fruits := [3]string{"apple", "orange", "banana"}
	fmt.Println(fruits)

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Example 5: Functions
	sum := add(3, 4)
	fmt.Println(sum)

	// Example 6: Structs
	person := Person{name: "Alice", age: 25}
	fmt.Println(person)

	// Example 7: Goroutines and Channels
	ch := make(chan string)
	go sendMessage(ch, "Hello, Go!")
	message := <-ch
	fmt.Println(message)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	name string
	age  int
}

func sendMessage(ch chan<- string, message string) {
	ch <- message
}