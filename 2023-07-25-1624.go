package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成随机数切片
	numbers := generateRandomNumbers(10)

	// 打印切片中的数字
	fmt.Println(numbers)
}

// 生成指定长度的随机数切片
func generateRandomNumbers(length int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}
