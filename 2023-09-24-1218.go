package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := generateNumbers(10)
	fmt.Println(numbers)

	average := calculateAverage(numbers)
	fmt.Println("Average:", average)
}

func generateNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}