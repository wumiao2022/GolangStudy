package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成10个随机数，并求平均值
	var sum int
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		sum += num
		fmt.Printf("第 %d 个随机数: %d\n", i+1, num)
	}
	average := float64(sum) / 10
	fmt.Printf("平均值: %.2f\n", average)
}