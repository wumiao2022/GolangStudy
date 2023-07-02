package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables and Constants
	var name string = "John"
	var age int = 25
	const pi float64 = 3.14159

	// Exercise 3: Basic Arithmetic Operations
	a, b := 10, 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	// Exercise 4: If-Else Statements
	if age >= 18 {
		fmt.Println(name, "is an adult.")
	} else {
		fmt.Println(name, "is a minor.")
	}

	// Exercise 5: For Loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}

	// Exercise 6: Arrays and Slices
	numbers := [5]int{1, 2, 3, 4, 5}
	slice := numbers[1:4]

	// Exercise 7: Functions
	result := add(3, 5)

	// Exercise 8: Pointers
	a := 10
	changeValue(&a)
	fmt.Println(a)

	// Exercise 9: Structures
	student := Student{Name: "Alice", Age: 20, Grade: "A"}

	// Exercise 10: Error Handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func add(a, b int) int {
	return a + b
}

func changeValue(a *int) {
	*a = 20
}

type Student struct {
	Name  string
	Age   int
	Grade string
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero.")
	} else {
		return a / b, nil
	}
}