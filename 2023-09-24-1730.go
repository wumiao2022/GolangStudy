package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 随机生成一个10个元素的整数数组
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	// 打印数组
	fmt.Println("原始数组:", arr)

	// 找到数组中的最大值和最小值并打印
	max, min := arr[0], arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("最大值:", max)
	fmt.Println("最小值:", min)
}