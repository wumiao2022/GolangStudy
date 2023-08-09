package main

import "fmt"

func main() {
	// Exercise 1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 3: Check if a number is even or odd
	number := 25
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// Exercise 4: Create a slice and iterate over it
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 5: Find the maximum element in an array
	array := []int{5, 9, 3, 1, 7}
	max := array[0]
	for _, num := range array {
		if num > max {
			max = num
		}
	}
	fmt.Println("The maximum element is", max)
}