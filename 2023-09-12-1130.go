package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Exercise 2: Print the numbers 1 to 10 using a loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 3: Calculate the sum of two numbers and print the result
	num1 := 3
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// Exercise 4: Calculate the factorial of a number using recursion
	num := 5
	fact := factorial(num)
	fmt.Println(fact)

	// Exercise 5: Calculate the average of a slice of numbers
	numbers := []float64{2.5, 3.6, 4.8, 6.2, 7.4}
	avg := calculateAverage(numbers)
	fmt.Println(avg)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	avg := sum / float64(len(numbers))
	return avg
}