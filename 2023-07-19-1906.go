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

	sum := sumNumbers(numbers)
	fmt.Println("Sum of numbers:", sum)
}

func generateRandomNumbers(count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func sumNumbers(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}