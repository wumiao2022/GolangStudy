package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个10个元素的随机整数数组
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("随机整数数组:", arr)

	// 计算数组元素的和
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("数组元素的和:", sum)

	// 找出数组中的最大值和最小值
	max := arr[0]
	min := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("数组中的最大值:", max)
	fmt.Println("数组中的最小值:", min)
}
