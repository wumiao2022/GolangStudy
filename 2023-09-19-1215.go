package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	nums := generateRandomNumbers(10)
	fmt.Println(nums)

	result := sumEvenNumbers(nums)
	fmt.Println(result)
}

func generateRandomNumbers(n int) []int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(100))
	}
	return nums
}

func sumEvenNumbers(nums []int) int {
	sum := 0
	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}