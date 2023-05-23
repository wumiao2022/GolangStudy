package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Hello world
	fmt.Println("Hello, world!")

	// Exercise 2: Variables and basic arithmetic
	a := 5
	b := 7
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)

	// Exercise 3: Conditional statements
	x := 10
	if x < 5 {
		fmt.Println("x is less than 5")
	} else if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is between 5 and 10")
	}

	// Exercise 4: Arrays and for loops
	nums := [5]int{2, 4, 6, 8, 10}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("The sum of the array is:", sum)

	// Exercise 5: Functions and return values
	res := add(3, 5)
	fmt.Println("The sum of 3 and 5 is:", res)
}

func add(a, b int) int {
	return a + b
}