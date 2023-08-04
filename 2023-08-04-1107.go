package main

import (
	"fmt"
	"strings"
)

func main() {
	// Exercise 1: Concatenate two strings with a space in between
	str1 := "Hello"
	str2 := "world"
	result := str1 + " " + str2
	fmt.Println(result)

	// Exercise 2: Convert a string to uppercase
	str3 := "hello"
	upper := strings.ToUpper(str3)
	fmt.Println(upper)

	// Exercise 3: Print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 4: Calculate the sum of numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Exercise 5: Check if a number is even or odd
	num := 24
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercise 6: Find the maximum element in an array
	arr := []int{2, 5, 9, 3, 7}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max)
}