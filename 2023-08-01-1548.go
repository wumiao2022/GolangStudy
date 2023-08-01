package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成随机数切片
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(100)
	}

	// 打印切片中的每个数值
	fmt.Println("Numbers:")
	for _, num := range numbers {
		fmt.Println(num)
	}
}