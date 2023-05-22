package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机整数切片
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println("nums:", nums)

	// 冒泡排序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println("sorted nums:", nums)
}