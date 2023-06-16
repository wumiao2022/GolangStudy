package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机整数数组
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("原始数组：", arr)

	// 冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("排序后数组：", arr)
}