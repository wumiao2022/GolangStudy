package main

import (
	"fmt"
)

func main() {
	// 1. Hello, World!
	fmt.Println("Hello, World!")

	// 2. Print even numbers
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 3. Calculate sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 4. Swap two numbers
	a := 10
	b := 20
	a, b = b, a
	fmt.Println("a:", a, "b:", b)

	// 5. Check if a number is positive or negative
	num := -5
	if num >= 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}

	// 6. Iterate over a string and print each character
	str := "Golang"
	for _, char := range str {
		fmt.Println(string(char))
	}

	// 7. Find the maximum of three numbers
	num3 := 10
	num4 := 20
	num5 := 30
	max := num3
	if num4 > max {
		max = num4
	}
	if num5 > max {
		max = num5
	}
	fmt.Println("Maximum:", max)

	// 8. Reverse a string
	str2 := "Hello, World!"
	reversedStr := ""
	for i := len(str2) - 1; i >= 0; i-- {
		reversedStr += string(str2[i])
	}
	fmt.Println("Reversed string:", reversedStr)
}