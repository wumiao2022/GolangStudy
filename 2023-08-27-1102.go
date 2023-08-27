package main

import "fmt"

// Example 1: Hello World
func helloWorld() {
	fmt.Println("Hello, World!")
}

// Example 2: Variables
func variables() {
	var name string = "John"
	age := 25

	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}

// Example 3: Arithmetic Operators
func arithmeticOperators() {
	a := 10
	b := 5

	sum := a + b
	diff := a - b
	mul := a * b
	div := a / b
	mod := a % b

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", mul)
	fmt.Printf("Quotient: %d\n", div)
	fmt.Printf("Remainder: %d\n", mod)
}

// Example 4: Conditional Statements
func conditionalStatements() {
	a := 10
	b := 5

	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}
}

// Example 5: Loops
func loops() {
	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While loop
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop with break statement
	k := 1
	for {
		fmt.Println(k)
		k++
		if k > 5 {
			break
		}
	}
}

// Example 6: Functions
func add(a, b int) int {
	return a + b
}

func functions() {
	result := add(5, 3)
	fmt.Printf("Result: %d\n", result)
}

// Example 7: Arrays
func arrays() {
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
}

// Example 8: Slices
func slices() {
	numbers := []int{1, 2, 3, 4, 5}

	// Append element to slice
	numbers = append(numbers, 6)

	// Slice a portion of the slice
	sliced := numbers[1:4]

	fmt.Println(numbers)
	fmt.Println(sliced)
}

// Example 9: Maps
func maps() {
	details := make(map[string]string)
	details["name"] = "John"
	details["age"] = "25"
	details["country"] = "USA"

	fmt.Println(details)
}

// Example 10: Structs
type Person struct {
	Name    string
	Age     int
	Country string
}

func structs() {
	person := Person{Name: "John", Age: 25, Country: "USA"}

	fmt.Println(person)
}

// Main function
func main() {
	helloWorld()
	variables()
	arithmeticOperators()
	conditionalStatements()
	loops()
	functions()
	arrays()
	slices()
	maps()
	structs()
}