package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成一个长度为10的随机整数数组
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("原数组：", arr)

	// 冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("排序后数组：", arr)

	// 计算数组平均值
	sum := 0
	for _, v := range arr {
		sum += v
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Printf("数组平均值：%.2f\n", avg)

	// 查找元素在数组中的位置
	target := arr[rand.Intn(len(arr))]
	for i := range arr {
		if arr[i] == target {
			fmt.Printf("元素 %d 在数组中的位置为：%d\n", target, i)
			break
		}
	}
}