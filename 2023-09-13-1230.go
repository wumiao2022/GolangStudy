package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers
	a := 5
	b := 3
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// Exercise 3: Find the maximum number in an array
	numbers := []int{10, 4, 8, 2, 9, 6, 5}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("The maximum number is", max)

	// Exercise 4: Reverse a string
	str := "Hello, Golang!"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("Reversed string:", reversed)

	// Exercise 5: Check if a number is prime
	number := 31
	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(number, "is prime:", isPrime)
}