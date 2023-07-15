package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个10个元素的整数数组
	array := generateArray(10)
	fmt.Println("Original Array:", array)

	// 计算数组中所有元素的和
	sum := calculateSum(array)
	fmt.Println("Sum of Array:", sum)

	// 找出数组中的最大值
	max := findMax(array)
	fmt.Println("Max Value:", max)

	// 数组反转
	reverseArray(array)
	fmt.Println("Reversed Array:", array)
}

// 生成指定大小的随机整数数组
func generateArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(100) // 生成0到99的随机整数
	}
	return array
}

// 计算整数数组中元素的和
func calculateSum(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

// 找出整数数组中的最大值
func findMax(array []int) int {
	max := array[0]
	for _, num := range array {
		if num > max {
			max = num
		}
	}
	return max
}

// 反转整数数组
func reverseArray(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}