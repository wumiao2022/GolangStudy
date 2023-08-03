package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, world!"
	fmt.Println("Hello, world!")

	// Exercise 2: Calculate the sum of two given numbers and print the result
	num1 := 10
	num2 := 15
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Print numbers from 1 to 10 using a loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 4: Check if a given number is even or odd and print the result
	number := 23
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// Exercise 5: Print the Fibonacci sequence up to a given number
	limit := 20
	fibonacci := []int{0, 1}

	for i := 2; fibonacci[i-1]+fibonacci[i-2] <= limit; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}

	fmt.Println("Fibonacci sequence:", fibonacci)
}