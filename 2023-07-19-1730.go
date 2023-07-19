package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数切片
	numbers := make([]int, 5)
	for i := 0; i < 5; i++ {
		numbers[i] = rand.Intn(100)
	}

	// 计算切片中元素的平均值
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))

	// 打印切片和平均值
	fmt.Println("Numbers:", numbers)
	fmt.Printf("Average: %.2f\n", average)
}