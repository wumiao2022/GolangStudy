package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world! Today's exercise: ")

	// Exercise 1: Hello, Golang
	fmt.Println("Exercise 1: Hello, Golang")
	fmt.Println("Hello, Golang!")

	// Exercise 2: Print the current time
	fmt.Println("Exercise 2: Print the current time")
	currentTime := time.Now()
	fmt.Println("Current time is:", currentTime.Format("15:04:05"))

	// Exercise 3: Calculate the sum of numbers
	fmt.Println("Exercise 3: Calculate the sum of numbers")
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("The sum of numbers is:", sum)

	// Exercise 4: Find the maximum number in an array
	fmt.Println("Exercise 4: Find the maximum number in an array")
	array := []int{10, 5, 8, 15, 3}
	max := array[0]
	for _, num := range array {
		if num > max {
			max = num
		}
	}
	fmt.Println("The maximum number in the array is:", max)

	// Exercise 5: Reverse a string
	fmt.Println("Exercise 5: Reverse a string")
	str := "Hello, Golang!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed string:", reversedStr)
}

// More exercises can be added above this line.