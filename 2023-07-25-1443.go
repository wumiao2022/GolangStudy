package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var age int = 25
	const name string = "John"
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Example 3: Arrays and Slices
	var fruits [3]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "orange"
	fmt.Println("Fruits:", fruits)

	colors := []string{"red", "green", "blue"}
	fmt.Println("Colors:", colors)

	// Example 4: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 5: If-else
	number := 10
	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	// Example 6: Functions
	result := add(5, 3)
	fmt.Println("Result of add function:", result)

	// Example 7: Pointers
	num := 10
	fmt.Println("Value of num:", num)
	increment(&num)
	fmt.Println("Value of num after increment:", num)

	// Example 8: Goroutines
	go printHello()
	time.Sleep(time.Second)

	// Example 9: Channels
	ch := make(chan string)
	go sendMessage(ch, "Hello, Golang!")
	msg := <-ch
	fmt.Println("Received Message:", msg)
}

func add(a, b int) int {
	return a + b
}

func increment(num *int) {
	*num++
}

func printHello() {
	fmt.Println("Hello from goroutine!")
}

func sendMessage(ch chan<- string, message string) {
	ch <- message
}