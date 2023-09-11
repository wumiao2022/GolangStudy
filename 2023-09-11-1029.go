package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机的整数数组
	nums := generateRandomArray(10)

	// 打印数组
	fmt.Println("随机生成的数组是：", nums)

	// 计算数组中的最大值和最小值
	max := findMax(nums)
	min := findMin(nums)
	fmt.Println("数组中的最大值是：", max)
	fmt.Println("数组中的最小值是：", min)
}

// 生成一个指定长度的随机整数数组
func generateRandomArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

// 找出数组中的最大值
func findMax(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// 找出数组中的最小值
func findMin(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}