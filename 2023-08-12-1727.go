package main

import "fmt"

func main() {
	// Example 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Example 2: Calculate the sum of two numbers
	fmt.Println(5 + 7)

	// Example 3: Reverse a string
	text := "Golang is awesome!"
	reversed := ""
	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}
	fmt.Println(reversed)

	// Example 4: Check if a number is even or odd
	number := 10
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Example 5: Find the maximum number in an array
	numbers := []int{3, 6, 2, 9, 1}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
}