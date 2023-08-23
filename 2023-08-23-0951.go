package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	name := "John Doe"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Exercise 3: Basic calculations
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)

	// Exercise 4: Conditional statements
	x := 20
	y := 30
	if x > y {
		fmt.Println("x is greater than y")
	} else if y > x {
		fmt.Println("y is greater than x")
	} else {
		fmt.Println("x and y are equal")
	}

	// Exercise 5: Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 7: Slices
	numbersSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(numbersSlice)

	// Exercise 8: Functions
	result := add(10, 20)
	fmt.Printf("Result: %d\n", result)

	// Exercise 9: Pointers
	num := 10
	fmt.Printf("Before: %d\n", num)
	modifyValue(&num)
	fmt.Printf("After: %d\n", num)

	// Exercise 10: Structs
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}

func modifyValue(num *int) {
	*num = 20
}

type Person struct {
	Name string
	Age  int
}