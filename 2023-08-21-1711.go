package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个长度为10的整数数组
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(100)
	}

	// 打印数组
	fmt.Println("随机生成的整数数组：", nums)

	// 计算数组中的最大值
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	// 打印最大值
	fmt.Println("数组中的最大值为：", max)
}