package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个长度为10的整数数组
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	// 打印数组
	fmt.Printf("生成的随机数组：%v\n", arr)

	// 计算数组总和
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Printf("数组总和：%d\n", sum)

	// 计算数组平均值
	average := float64(sum) / float64(len(arr))
	fmt.Printf("数组平均值：%.2f\n", average)

	// 查找数组中的最大值和最小值
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
	fmt.Printf("数组最大值：%d\n", max)
	fmt.Printf("数组最小值：%d\n", min)
}