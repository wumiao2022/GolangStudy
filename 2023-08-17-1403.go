package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var name string = "John"
	var age int = 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Exercise 3: Data Types
	var number int = 10
	var isExist bool = true
	var pi float64 = 3.14
	fmt.Println(number, isExist, pi)

	// Exercise 4: If-Else statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is equal to 5")
	}

	// Exercise 5: For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Array
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 7: Slice
	var fruits []string = []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// Exercise 8: Map
	var person map[string]string = map[string]string{"name": "John", "age": "25"}
	fmt.Println(person)

	// Exercise 9: Function
	greet("Alice")

	// Exercise 10: Struct
	type Person struct {
		name string
		age  int
	}

	p := Person{name: "Bob", age: 30}
	fmt.Println(p)
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}