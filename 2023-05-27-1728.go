package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 10)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}

	fmt.Println("Original Numbers:", numbers)

	selectionSort(numbers)

	fmt.Println("Sorted Numbers:", numbers)
}

func selectionSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
		}
	}
}