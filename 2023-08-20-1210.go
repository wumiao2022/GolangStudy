package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数切片
	numbers := generateRandomNumbers(10)

	// 打印切片中的元素
	fmt.Println("随机整数切片:", numbers)

	// 计算切片中的元素之和
	sum := calculateSum(numbers)
	fmt.Println("元素之和:", sum)

	// 找到切片中的最大元素
	max := findMax(numbers)
	fmt.Println("最大元素:", max)

	// 找到切片中的最小元素
	min := findMin(numbers)
	fmt.Println("最小元素:", min)
}

// 生成一个包含n个随机整数的切片
func generateRandomNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

// 计算切片中的元素之和
func calculateSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// 找到切片中的最大元素
func findMax(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// 找到切片中的最小元素
func findMin(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}