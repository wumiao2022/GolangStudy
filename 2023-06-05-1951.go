package main

import "fmt"

func main() {
	// Example 1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Add two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// Example 3: Calculate the average of three numbers
	num3 := 30
	avg := (num1 + num2 + num3) / 3.0
	fmt.Printf("The average of %d, %d and %d is %f\n", num1, num2, num3, avg)

	// Example 4: check if a number is even or odd
	num4 := 15
	if num4%2 == 0 {
		fmt.Printf("%d is even\n", num4)
	} else {
		fmt.Printf("%d is odd\n", num4)
	}

	// Example 5: Find the maximum number in an array
	numbers := []int{10, 20, 30, 40, 50}
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Printf("The maximum number in the array %v is %d\n", numbers, max)
}