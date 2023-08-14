package main

import "fmt"

func main() {
	// Example 1: Print 'Hello, World!'
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 3: Swap two numbers
	a := 5
	b := 10
	fmt.Println("Before swap: a =", a, ", b =", b)
	temp := a
	a = b
	b = temp
	fmt.Println("After swap: a =", a, ", b =", b)

	// Example 4: Check if a number is even or odd
	number := 7
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// Example 5: Find the maximum number in an array
	numbers := []int{5, 2, 9, 7, 3}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Maximum number:", max)
}