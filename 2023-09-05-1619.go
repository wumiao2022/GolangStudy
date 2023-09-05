package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of 2 numbers and print the result
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Exercise 3: Create a function that takes in an integer parameter and returns its square
	number := 7
	square := getSquare(number)
	fmt.Println("Square of", number, "is", square)

	// Exercise 4: Create a struct called "Person" with name and age fields, and then print a person's information
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "John", Age: 25}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// Exercise 5: Generate a fibonacci sequence and print the first 10 numbers
	fmt.Println("Fibonacci Series:")
	fibonacci := generateFibonacci(10)
	fmt.Println(fibonacci)
}

func getSquare(num int) int {
	return num * num
}

func generateFibonacci(n int) []int {
	fibonacci := []int{0, 1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	return fibonacci
}