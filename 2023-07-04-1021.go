package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 多练习 - 生成并打印一个包含10个随机整数的切片
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(100)
	}
	fmt.Println(numbers)
}