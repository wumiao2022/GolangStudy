package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello World!")

	// Example 2: Variables and Constants
	var name string = "John"
	age := 30
	const pi float64 = 3.14159

	fmt.Printf("Name: %v\nAge: %v\nPi: %v\n", name, age, pi)

	// Example 3: Data Types
	var a int = 10
	var b float64 = 10.5
	var c string = "Hello"
	var d bool = true

	fmt.Printf("a: %d\nb: %f\nc: %s\nd: %t\n", a, b, c, d)

	// Example 4: Arrays and Slices
	var array1 [3]int = [3]int{1, 2, 3}
	var slice1 []int = []int{4, 5, 6}

	slice1 = append(slice1, 7)

	fmt.Printf("Array1: %v\nSlice1: %v\n", array1, slice1)

	// Example 5: Loops and Conditions
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even.\n", i)
		} else {
			fmt.Printf("%d is odd.\n", i)
		}
	}

	// Example 6: Functions
	num1 := 10
	num2 := 20

	fmt.Printf("Sum of %d and %d is %d.\n", num1, num2, sum(num1, num2))

	// Example 7: Pointers
	num := 10
	increment(&num)

	fmt.Printf("Num after increment by 1 is %d.\n", num)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func increment(num *int) {
	*num += 1
}