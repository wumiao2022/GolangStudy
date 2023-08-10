package main

import "fmt"

func main() {
	// Example 1: Calculate the sum of two numbers
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// Example 2: Count the number of vowels in a string
	str := "Hello World"
	vowelCount := 0
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, char := range str {
		for _, vowel := range vowels {
			if char == vowel || char == vowel-32 {
				vowelCount++
			}
		}
	}
	fmt.Println("Vowel count:", vowelCount)

	// Example 3: Check if a number is a prime number
	num := 23
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// Example 4: Reverse a string
	str = "Hello World"
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	fmt.Println("Reversed string:", reversedStr)
}