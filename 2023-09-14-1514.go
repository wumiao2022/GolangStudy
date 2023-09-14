package main

import "fmt"

func main() {
	// 1.
	fmt.Println("Hello, World!")

	// 2.
	var num1 int = 10
	var num2 int = 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3.
	var isTrue bool = true
	if isTrue {
		fmt.Println("It's true!")
	} else {
		fmt.Println("It's false!")
	}

	// 4.
	var fullName string = "John Doe"
	fmt.Println("Full Name:", fullName)

	// 5.
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println(number)
	}
}