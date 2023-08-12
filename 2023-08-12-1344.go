package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数数组
	numbers := make([]int, 10)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}

	// 打印数组内容
	fmt.Println("随机数数组：", numbers)

	// 计算数组元素和
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	// 打印数组元素和
	fmt.Println("数组元素和：", sum)
}