package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 实现一个函数，生成一个指定长度的随机整数数组，并打印出数组内容和长度
	printRandomArray(10)
}

func printRandomArray(length int) {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Printf("随机整数数组长度：%d，内容：%v\n", length, arr)
}