package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成随机数切片
	nums := make([]int, 5)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}

	// 打印随机数切片
	fmt.Println("随机数切片: ", nums)

	// 计算随机数的和
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 打印随机数的和
	fmt.Println("随机数的和: ", sum)
}