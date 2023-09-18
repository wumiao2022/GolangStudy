package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数切片
	nums := generateRandomSlice(10)

	// 计算切片中所有元素的平均值
	avg := calculateAverage(nums)

	// 输出结果
	fmt.Printf("随机整数切片: %v\n", nums)
	fmt.Printf("平均值: %.2f\n", avg)
}

// 生成指定长度的随机整数切片
func generateRandomSlice(length int) []int {
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

// 计算整数切片中所有元素的平均值
func calculateAverage(nums []int) float64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}