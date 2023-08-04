package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个数组
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	// 打印数组
	fmt.Println("原始数组：", arr)

	// 对数组进行升序排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	// 打印排序后的数组
	fmt.Println("升序排序后的数组：", arr)
}
