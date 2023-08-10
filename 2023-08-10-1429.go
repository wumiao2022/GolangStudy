package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Today's date:", currentTime.Format("2006-01-02"))

	// Exercise 1: Fibonacci Series
	n := 10
	fmt.Println("Fibonacci Series up to", n, "terms:")
	fibonacciSeries(n)

	// Exercise 2: Prime Numbers
	maxNum := 20
	fmt.Println("Prime numbers up to", maxNum, ":")
	primeNumbers(maxNum)

	// Exercise 3: Odd or Even Sum
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := oddEvenSum(arr)
	fmt.Println("Sum of Odd numbers:", sum[0])
	fmt.Println("Sum of Even numbers:", sum[1])

	// Exercise 4: Palindrome Check
	word := "racecar"
	isPalindrome := palindromeCheck(word)
	if isPalindrome {
		fmt.Println("The word", word, "is a palindrome.")
	} else {
		fmt.Println("The word", word, "is not a palindrome.")
	}
}

// Fibonacci Series
func fibonacciSeries(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

// Prime Numbers
func primeNumbers(maxNum int) {
	for i := 2; i <= maxNum; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// Odd or Even Sum
func oddEvenSum(arr []int) [2]int {
	sum := [2]int{0, 0}
	for _, num := range arr {
		if num%2 == 0 {
			sum[1] += num // Even sum
		} else {
			sum[0] += num // Odd sum
		}
	}
	return sum
}

// Palindrome Check
func palindromeCheck(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}