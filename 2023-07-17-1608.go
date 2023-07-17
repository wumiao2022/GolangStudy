package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1: 生成一个随机整数数组，并输出数组中的所有元素
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println(nums)

	// 练习2: 计算数组中所有元素的和，并输出结果
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// 练习3: 找出数组中的最大值，并输出结果
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	fmt.Println("Max:", max)
}