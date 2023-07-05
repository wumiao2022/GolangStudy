package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	numbers := make([]int, 10)

	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(100)
	}

	fmt.Println("Unsorted numbers:", numbers)

	bubbleSort(numbers)

	fmt.Println("Sorted numbers:", numbers)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}