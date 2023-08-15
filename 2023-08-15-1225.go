package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成一个随机数切片
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 0)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	// 对切片进行排序
	Sort(numbers)

	// 打印排序结果
	fmt.Println(numbers)
}

// 冒泡排序
func Sort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}
