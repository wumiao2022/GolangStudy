package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := generateRandomNumbers(10)
	fmt.Println("Generated random numbers:", nums)
}

func generateRandomNumbers(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}