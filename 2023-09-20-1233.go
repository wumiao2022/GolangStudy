package main

import "fmt"

// Exercise 1
func printHelloWorld() {
	fmt.Println("Hello, World!")
}

// Exercise 2
func calculateSum(a, b int) int {
	return a + b
}

// Exercise 3
func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

// Exercise 4
func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

// Exercise 5
func fibonacciSequence(n int) {
	f0 := 0
	f1 := 1
	fmt.Print(f0, " ", f1)

	for i := 2; i < n; i++ {
		f2 := f0 + f1
		fmt.Print(" ", f2)
		f0 = f1
		f1 = f2
	}
	fmt.Println()
}

func main() {
	// Exercise 1: Print "Hello, World!"
	printHelloWorld()

	// Exercise 2: Calculate and print the sum of two integers
	sum := calculateSum(5, 7)
	fmt.Println("Sum:", sum)

	// Exercise 3: Check if a number is even or odd
	checkEvenOdd(10)
	checkEvenOdd(15)

	// Exercise 4: Reverse a string
	fmt.Println(reverseString("Hello, World!"))

	// Exercise 5: Print the Fibonacci sequence
	fibonacciSequence(10)
}