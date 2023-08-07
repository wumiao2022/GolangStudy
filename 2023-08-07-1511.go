package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Print the sum of two numbers
	num1 := 20
	num2 := 10
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Exercise 3: Print the average of three numbers
	num3 := 6
	num4 := 8
	num5 := 10
	average := (num3 + num4 + num5) / 3
	fmt.Println("The average of", num3, ",", num4, "and", num5, "is", average)

	// Exercise 4: Check if a number is even or odd
	checkNumber := 15
	if checkNumber%2 == 0 {
		fmt.Println(checkNumber, "is even")
	} else {
		fmt.Println(checkNumber, "is odd")
	}

	// Exercise 5: Print a pattern using nested loops
	patternSize := 5
	for i := 1; i <= patternSize; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Exercise 6: Find the maximum number in an array
	numbers := []int{10, 4, 8, 6, 12, 2}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("The maximum number in the array is", max)
}
