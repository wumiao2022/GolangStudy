package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Calculate the average of three numbers
	num3 := 30
	average := (num1 + num2 + num3) / 3
	fmt.Println("Average:", average)

	// Exercise 4: Print the even numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Exercise 5: Reverse a string
	str := "Hello, Go!"
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	fmt.Println("Reversed String:", reversedStr)
}