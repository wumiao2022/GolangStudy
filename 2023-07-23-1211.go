package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Print variable
	num := 25
	fmt.Println("The value of num is:", num)

	// Exercise 3: Calculate average
	num1 := 10
	num2 := 20
	num3 := 30
	average := (num1 + num2 + num3) / 3
	fmt.Println("The average is:", average)

	// Exercise 4: Conditional statement
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a child.")
	}

	// Exercise 5: Loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}