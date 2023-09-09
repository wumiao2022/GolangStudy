package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Hello World
	fmt.Println("Hello, World!")

	// 2. Variables
	var name string = "Alice"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 3. Constants
	const pi = 3.14159
	fmt.Println("The value of pi is", pi)

	// 4. Input and Output
	var input string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&input)
	fmt.Println("Hello,", input)

	// 5. Control Structures
	num := 10
	if num > 5 {
		fmt.Println("The number is greater than 5.")
	} else {
		fmt.Println("The number is less than or equal to 5.")
	}

	// 6. Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 7. Arrays and Slices
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Array:", array)
	fmt.Println("Slice:", slice)

	// 8. Functions
	fmt.Println("Sum:", sum(3, 5))

	// 9. Pointers
	numPtr := &num
	fmt.Println("Value:", num)
	fmt.Println("Address:", numPtr)

	// 10. Structs
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Bob", Age: 30}
	fmt.Println("Person:", person)

	// 11. Packages
	fmt.Println("Today is", time.Now().Weekday())
}

func sum(a, b int) int {
	return a + b
}