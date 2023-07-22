package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := generateNumbers(10)
	fmt.Println("Generated numbers:", numbers)
	highest := findHighest(numbers)
	fmt.Println("Highest number:", highest)
}

func generateNumbers(count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func findHighest(numbers []int) int {
	highest := numbers[0]
	for _, num := range numbers {
		if num > highest {
			highest = num
		}
	}
	return highest
}