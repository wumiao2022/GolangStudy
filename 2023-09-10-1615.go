package main

import "fmt"

func main() {
	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Fibonacci Series
	fibonacci := func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}

	// Reverse a string
	reverseString := func(str string) string {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	fmt.Println(reverseString("Hello, World!"))
}