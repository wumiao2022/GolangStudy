package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// Exercise 1: Print numbers from 1-10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 2: Print "Golang" 10 times
	for i := 0; i < 10; i++ {
		fmt.Println("Golang")
	}

	// Exercise 3: Calculate the sum of numbers from 1-10 using a loop
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Exercise 4: Print even numbers from 1-20
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 5: Print a pattern of asterisks
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}