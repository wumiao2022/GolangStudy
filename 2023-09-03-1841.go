package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算1+2的结果并打印
	sum := 1 + 2
	fmt.Println(sum)

	// 练习3: 使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4: 使用if条件判断语句判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("是奇数")
	}

	// 练习5: 定义一个函数add，接收两个参数并返回它们的和
	result := add(3, 4)
	fmt.Println(result)
}

// add函数，接收两个参数a和b，返回它们的和
func add(a, b int) int {
	return a + b
}
