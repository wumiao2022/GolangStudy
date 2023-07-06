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

	// 打印切片中的所有元素
	printSlice(numbers)

	// 计算切片中的最小值和最大值
	min, max := findMinMax(numbers)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// 计算切片中的总和
	sum := calculateSum(numbers)
	fmt.Printf("Sum: %d\n", sum)
}

// 生成指定长度的随机整数切片
func generateRandomNumbers(length int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

// 打印整数切片中的所有元素
func printSlice(numbers []int) {
	fmt.Println("Numbers in the slice:")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// 找到整数切片中的最小值和最大值
func findMinMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

// 计算整数切片中的总和
func calculateSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}