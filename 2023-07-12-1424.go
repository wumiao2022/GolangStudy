package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Calculate the sum of two numbers and print the result
	a := 3
	b := 7
	sum := a + b
	fmt.Println("Sum:", sum)

	// Exercise 3: Check if a given number is even or odd and print the result
	num := 14
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercise 4: Find the largest number in a given slice and print it
	numbers := []int{5, 2, 9, 1, 7}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Largest number:", max)

	// Exercise 5: Generate and print Fibonacci sequence up to a given limit
	limit := 10
	fmt.Println("Fibonacci sequence:")
	fibonacci := []int{0, 1}
	for i := 2; i < limit; i++ {
		next := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, next)
	}
	fmt.Println(fibonacci)
}
```