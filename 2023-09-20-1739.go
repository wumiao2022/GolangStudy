package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Swap the values of two variables
	a := 5
	b := 10
	a, b = b, a
	fmt.Println("a =", a, "b =", b)

	// Exercise 4: Check if a number is even or odd
	number := 9
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercise 5: Find the largest number in an array
	numbers := []int{5, 2, 9, 10, 4}
	largest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
	}
	fmt.Println("Largest number:", largest)
}