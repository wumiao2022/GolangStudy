package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numbers := generateRandomNumbers(5)
	fmt.Println("Random numbers:", numbers)

	sortedNumbers := bubbleSort(numbers)
	fmt.Println("Sorted numbers:", sortedNumbers)
}

func generateRandomNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func bubbleSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}
