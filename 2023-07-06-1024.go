package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：输出当前时间
	fmt.Println(time.Now())

	// 练习3：计算两个整数的和并输出
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习4：遍历并输出切片元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习5：创建和输出一个字典
	dict := make(map[string]int)
	dict["apple"] = 10
	dict["orange"] = 15
	dict["banana"] = 8

	for key, value := range dict {
		fmt.Println(key, value)
	}
}

// 每日多练，继续加油！