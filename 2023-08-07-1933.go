package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numbers := generateRandomNumbers(10)
	fmt.Println("Random numbers:", numbers)

	sum := calculateSum(numbers)
	fmt.Println("Sum:", sum)

	average := calculateAverage(numbers)
	fmt.Println("Average:", average)

	max := findMax(numbers)
	fmt.Println("Max value:", max)

	min := findMin(numbers)
	fmt.Println("Min value:", min)
}

func generateRandomNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func calculateSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateAverage(numbers []int) float64 {
	sum := calculateSum(numbers)
	average := float64(sum) / float64(len(numbers))
	return average
}

func findMax(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func findMin(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}