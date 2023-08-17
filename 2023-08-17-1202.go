package main

import "fmt"

func main() {
	// Exercise 1: Hello, World!
	fmt.Println("Hello, World!")

	// Exercise 2: Print odd numbers from 1 to 10
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// Exercise 3: Calculate the sum of numbers from 1 to 100
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// Exercise 4: Print the Fibonacci series up to 10 terms
	n := 10
	first, second := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(first)
		temp := first + second
		first = second
		second = temp
	}

	// Exercise 5: Reverse a string
	str := "Hello, World!"
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	fmt.Println("Reverse:", reverse)
}

// Exercise 6: Write a function to check if a given number is prime
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Exercise 7: Write a function to find the maximum element in an array
func findMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// Exercise 8: Implement bubble sort algorithm to sort an array in ascending order
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}