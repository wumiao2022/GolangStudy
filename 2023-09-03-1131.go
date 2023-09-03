package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1: 生成一个随机整数数组，长度为10，并打印出数组元素
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	// 练习2: 使用冒泡排序对数组进行升序排序，并打印排序后的数组
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

	// 练习3: 计算数组中所有元素的和，并打印结果
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println(sum)

	// 练习4: 查找数组中的最大值，并打印结果
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
}
