package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Perform basic arithmetic operations
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)
	fmt.Printf("Remainder: %d\n", remainder)

	// Exercise 3: Print numbers from 1 to 10 using a loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Exercise 4: Calculate the factorial of a given number
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d: %d\n", num, factorial)

	// Exercise 5: Find the maximum number from an array
	numbers := []int{5, 2, 9, 1, 7}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Printf("Max number: %d\n", max)
}