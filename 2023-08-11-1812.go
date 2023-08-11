package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables
	var name string = "John"
	age := 30
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Example 3: Constants
	const pi = 3.14159
	const gravity = 9.8
	fmt.Printf("The value of pi is %.2f and the acceleration due to gravity is %.2f m/s^2.\n", pi, gravity)

	// Example 4: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Example 5: Slices
	slice := numbers[1:4]
	fmt.Println("Slice:", slice)

	// Example 6: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 7: Conditionals
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is smaller than or equal to 5")
	}

	// Example 8: Functions
	add := func(a, b int) int {
		return a + b
	}
	result := add(2, 3)
	fmt.Println("Result of addition:", result)

	// Example 9: Structures
	type person struct {
		name string
		age  int
	}
	p := person{name: "Alice", age: 25}
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)

	// Example 10: Pointers
	y := 10
	yPtr := &y
	fmt.Println("Value of y:", *yPtr)

	// Example 11: Goroutines
	go hello()

	// Example 12: Channels
	message := make(chan string)
	go process(message)
	msg := <-message
	fmt.Println("Received message:", msg)
}

func hello() {
	fmt.Println("Hello, Goroutine!")
}

func process(ch chan<- string) {
	ch <- "Message from Goroutine"
}