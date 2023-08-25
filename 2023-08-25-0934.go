package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数数组
	arr := generateRandomArray(10)

	// 打印数组
	fmt.Println("Original array:", arr)

	// 计算数组的最大值
	max := findMax(arr)
	fmt.Println("Max value:", max)

	// 数组反转
	reverseArray(arr)
	fmt.Println("Reversed array:", arr)

	// 数组元素乘2
	multiplyArray(arr, 2)
	fmt.Println("Multiplied array:", arr)

	// 数组元素相加
	addArray(arr, 10)
	fmt.Println("Added array:", arr)
}

// 生成随机数数组
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

// 查找数组最大值
func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// 数组反转
func reverseArray(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 数组元素乘以n
func multiplyArray(arr []int, n int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * n
	}
}

// 数组元素加上n
func addArray(arr []int, n int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + n
	}
}
