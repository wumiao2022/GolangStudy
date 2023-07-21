package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variable declaration and printing
	var name string = "John"
	var age int = 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Exercise 3: Conditional statements
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// Exercise 4: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Arrays and loops
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 6: Functions
	result := multiply(3, 4)
	fmt.Println(result)
}

func multiply(a, b int) int {
	return a * b
}