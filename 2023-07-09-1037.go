package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 3: Calculate the sum of two numbers
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 4: Find the maximum number in an array
	numbers := []int{2, 7, 9, 4, 1, 5}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("The maximum number is", max)

	// Exercise 5: Print a pattern
	size := 5
	for i := 1; i <= size; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}