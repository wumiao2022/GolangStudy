package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个长度为10的随机整数数组
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	// 打印数组
	fmt.Println("Original array:", arr)

	// 找到数组中的最大值和最小值
	min := arr[0]
	max := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// 打印最大值和最小值
	fmt.Println("Minimum value:", min)
	fmt.Println("Maximum value:", max)
}
