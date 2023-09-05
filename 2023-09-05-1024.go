package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Print Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Print Numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 3: Print Even Numbers from 1 to 20
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 4: Calculate the Sum of Numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Exercise 5: Find the Largest Number in an Array
	numbers := []int{10, 5, 2, 15, 20}
	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	fmt.Println("Largest Number:", largest)
}