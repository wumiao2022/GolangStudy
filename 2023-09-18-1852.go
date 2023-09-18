package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1: 生成一个随机整数切片，长度为10，元素范围在0-100之间
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(101)
	}
	fmt.Println("随机整数切片:", slice)

	// 练习2: 计算切片中所有元素之和
	sum := 0
	for _, num := range slice {
		sum += num
	}
	fmt.Println("切片元素之和:", sum)

	// 练习3: 找出切片中的最大值和最小值，并输出
	max := slice[0]
	min := slice[0]
	for _, num := range slice {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("最大值:", max)
	fmt.Println("最小值:", min)
}