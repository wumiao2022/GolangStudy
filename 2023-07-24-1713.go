package main

import "fmt"

// Exercise 1: Print "Hello, World!"
func exercise1() {
	fmt.Println("Hello, World!")
}

// Exercise 2: Sum of two numbers
func exercise2(num1, num2 int) int {
	return num1 + num2
}

// Exercise 3: Reverse a string
func exercise3(str string) string {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// Exercise 4: Check if a number is even or odd
func exercise4(num int) string {
	if num%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// Exercise 5: Find the maximum number in an array
func exercise5(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// Main function to run the exercises
func main() {
	exercise1()
	fmt.Println(exercise2(3, 5))
	fmt.Println(exercise3("Hello"))
	fmt.Println(exercise4(7))
	fmt.Println(exercise5([]int{1, 5, 3, 9, 2}))
}
