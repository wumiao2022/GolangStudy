package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和并输出结果
	var a, b int = 10, 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3: 使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4: 使用if-else语句判断一个数是奇数还是偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习5: 创建一个简单的函数来计算两个数的差
	fmt.Println(subtract(10, 5))
}

func subtract(a, b int) int {
	return a - b
}