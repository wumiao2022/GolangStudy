package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := []int{}
	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.Intn(100))
	}
	fmt.Println("Before sorting:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)
}

func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}