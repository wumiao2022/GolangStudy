package main

import (
	"fmt"
)

func main() {
	// Example 1
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)

	// Example 2
	var str string
	fmt.Print("Please input a string: ")
	fmt.Scan(&str)
	fmt.Printf("The input string is: %s\n", str)

	// Example 3
	nums := []int{10, 20, 30, 40, 50}
	for i, num := range nums {
		fmt.Printf("The value of element %d is: %d\n", i, num)
	}

	// Example 4
	count := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			count++
		}
	}
	fmt.Printf("There are %d even numbers between 1 and 100\n", count)

	// Example 5
	num := 5
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	fmt.Printf("The factorial of %d is: %d\n", num, result)
}