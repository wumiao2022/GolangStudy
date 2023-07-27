package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Printf("Today's date is: %s\n", currentTime.Format("02-Jan-2006"))

	// Exercise 1: Sum of integers
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Sum of numbers: %d\n", sum)

	// Exercise 2: Calculate factorial
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d: %d\n", num, factorial)

	// Exercise 3: Reverse a string
	str := "Hello, Go!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Printf("Reversed string: %s\n", reversedStr)
}