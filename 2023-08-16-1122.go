package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	fmt.Println(time.Now())

	// 使用循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 定义一个切片并进行遍历
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 定义一个函数，计算两个整数的和
	sum := add(5, 3)
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}