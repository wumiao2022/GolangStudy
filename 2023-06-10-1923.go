package main

import "fmt"

func main() {
	// Example 1: Simple Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variables and Constants
	var message string = "Hello, Golang!"
	const pi float64 = 3.14159

	fmt.Println(message)
	fmt.Println(pi)

	// Example 3: Basic Data Types
	var age int = 28
	var weight float64 = 65.4
	var isCoding bool = true
	var name string = "Alice"

	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(isCoding)
	fmt.Println(name)

	// Example 4: Basic Math
	var x int = 10
	var y int = 3

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// Example 5: Control Structures
	var n int = 5

	if n > 5 {
		fmt.Println("n is greater than 5")
	} else if n < 5 {
		fmt.Println("n is less than 5")
	} else {
		fmt.Println("n is equal to 5")
	}

	// Example 6: Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 7: Functions
	fmt.Println(getSum(1, 2))

	// Example 8: Arrays and Slices
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var slice []int = []int{1, 2, 3}

	fmt.Println(arr[0])
	fmt.Println(slice[1])

	// Example 9: Structs
	type Person struct {
		name string
		age  int
	}

	var person1 Person = Person{name: "Bob", age: 25}
	var person2 Person = Person{"Alice", 28}

	fmt.Println(person1.name, person1.age)
	fmt.Println(person2.name, person2.age)
}

func getSum(x, y int) int {
	return x + y
}