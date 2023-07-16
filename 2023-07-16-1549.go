package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个长度为10的随机整数数组，元素取值范围在1-100之间
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}

	// 输出数组元素
	fmt.Println("数组元素：", arr)

	// 计算数组元素的平均值
	sum := 0
	for _, num := range arr {
		sum += num
	}
	average := float64(sum) / float64(len(arr))
	fmt.Println("平均值：", average)

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
	fmt.Println("最大值：", max)
	fmt.Println("最小值：", min)
}