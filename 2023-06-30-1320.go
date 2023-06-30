package main

import "fmt"

func main() {
	// 1. Print "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. Calculate the sum of two numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. Print all even numbers between 1 and 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 4. Create a simple function to check if a number is odd or even
	checkOddEven := func(num int) {
		if num%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	checkOddEven(3) // Output: Odd
	checkOddEven(6) // Output: Even

	// 5. Create a slice of fruits and print each fruit using a loop
	fruits := []string{"apple", "banana", "orange"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}