package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数数组
	arr := generateRandomArray(10, 100)

	// 打印数组
	fmt.Println("原始数组：", arr)

	// 对数组进行排序
	bubbleSort(arr)

	// 打印排序后的数组
	fmt.Println("排序后的数组：", arr)
}

// 生成一个指定长度的随机整数数组，值的范围为[0, max)
func generateRandomArray(length, max int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}

// 冒泡排序
func bubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}